#include "src/carnot/compiler/objects/pl_module.h"
#include "src/carnot/compiler/objects/expr_object.h"
#include "src/carnot/compiler/objects/test_utils.h"
#include "src/shared/metadata/base_types.h"

namespace pl {
namespace carnot {
namespace compiler {
using ::pl::table_store::schema::Relation;

constexpr char kRegInfoProto[] = R"proto(
scalar_udfs {
  name: "pl.equals"
  exec_arg_types: UINT128
  exec_arg_types: UINT128
  return_type: BOOLEAN
}
)proto";

constexpr char kUDTFSourcePb[] = R"proto(
name: "OpenNetworkConnections"
args {
  name: "upid"
  arg_type: STRING
  semantic_type: ST_UPID
}
executor: UDTF_SUBSET_PEM
relation {
  columns {
    column_name: "time_"
    column_type: TIME64NS
  }
  columns {
    column_name: "fd"
    column_type: INT64
  }
  columns {
    column_name: "name"
    column_type: STRING
  }
}
)proto";

class PLModuleTest : public QLObjectTest {
 protected:
  std::unique_ptr<compiler::RegistryInfo> SetUpRegistryInfo() {
    udfspb::UDFInfo udf_proto;
    CHECK(google::protobuf::TextFormat::MergeFromString(kRegInfoProto, &udf_proto));

    auto info = std::make_unique<compiler::RegistryInfo>();
    PL_CHECK_OK(info->Init(udf_proto));
    udfspb::UDTFSourceSpec spec;
    google::protobuf::TextFormat::MergeFromString(kUDTFSourcePb, &spec);
    info->AddUDTF(spec);
    return info;
  }

  std::unique_ptr<RelationMap> SetUpRelMap() {
    auto rel_map = std::make_unique<RelationMap>();
    rel_map->emplace("sequences", Relation(
                                      {
                                          types::TIME64NS,
                                          types::FLOAT64,
                                          types::FLOAT64,
                                      },
                                      {"time_", "xmod10", "PIx"}));
    return rel_map;
  }

  void SetUp() override {
    QLObjectTest::SetUp();
    info_ = SetUpRegistryInfo();
    compiler_state_ = std::make_unique<CompilerState>(SetUpRelMap(), info_.get(), time_now_);
    module_ = PLModule::Create(graph.get(), compiler_state_.get()).ConsumeValueOrDie();
  }

  std::unique_ptr<CompilerState> compiler_state_;
  int64_t time_now_ = 1552607213931245000;
  std::unique_ptr<RegistryInfo> info_;
  std::shared_ptr<PLModule> module_;
};

TEST_F(PLModuleTest, ModuleFindAttributeFromRegistryInfo) {
  auto attr_or_s = module_->GetAttribute(ast, "equals");

  ASSERT_OK(attr_or_s);
  QLObjectPtr attr_object = attr_or_s.ConsumeValueOrDie();

  ASSERT_FALSE(attr_object->HasNode());
  ASSERT_TRUE(attr_object->type_descriptor().type() == QLObjectType::kFunction);
  auto result_or_s =
      std::static_pointer_cast<FuncObject>(attr_object)->Call({}, ast, ast_visitor.get());
  ASSERT_OK(result_or_s);
  auto ql_object = result_or_s.ConsumeValueOrDie();
  ASSERT_TRUE(Match(ql_object->node(), Func()));

  FuncIR* func = static_cast<FuncIR*>(ql_object->node());
  EXPECT_EQ(func->carnot_op_name(), "equals");
}

TEST_F(PLModuleTest, AttributeNotFound) {
  std::string attribute = "bar";
  auto attr_or_s = module_->GetAttribute(ast, attribute);

  ASSERT_NOT_OK(attr_or_s);
  EXPECT_THAT(attr_or_s.status(), HasCompilerError("'pl' object has no attribute .*$0", attribute));
}

TEST_F(PLModuleTest, GetUDTFMethod) {
  std::string upid_value = "11285cdd-1de9-4ab1-ae6a-0ba08c8c676c";
  auto upid_str = MakeString(upid_value);
  std::string network_conns_udtf_name = "OpenNetworkConnections";
  auto method_or_s = module_->GetMethod(network_conns_udtf_name);

  ASSERT_OK(method_or_s);
  QLObjectPtr method_object = method_or_s.ConsumeValueOrDie();

  ASSERT_TRUE(method_object->type_descriptor().type() == QLObjectType::kFunction);
  auto result_or_s = std::static_pointer_cast<FuncObject>(method_object)
                         ->Call({{{"upid", upid_str}}, {}}, ast, ast_visitor.get());
  ASSERT_OK(result_or_s);
  auto ql_object = result_or_s.ConsumeValueOrDie();
  ASSERT_TRUE(ql_object->type_descriptor().type() == QLObjectType::kDataframe);
  ASSERT_TRUE(Match(ql_object->node(), UDTFSource()));

  auto udtf = static_cast<UDTFSourceIR*>(ql_object->node());
  EXPECT_EQ(udtf->func_name(), network_conns_udtf_name);
  const auto& arg_values = udtf->arg_values();
  ASSERT_EQ(arg_values.size(), 1);
  EXPECT_EQ(upid_str, arg_values[0]);
}

TEST_F(PLModuleTest, GetUDTFMethodBadArguements) {
  std::string network_conns_udtf_name = "OpenNetworkConnections";
  auto method_or_s = module_->GetMethod(network_conns_udtf_name);

  ASSERT_OK(method_or_s);
  QLObjectPtr method_object = method_or_s.ConsumeValueOrDie();

  ASSERT_TRUE(method_object->type_descriptor().type() == QLObjectType::kFunction);
  auto result_or_s =
      std::static_pointer_cast<FuncObject>(method_object)->Call({}, ast, ast_visitor.get());
  ASSERT_NOT_OK(result_or_s);
  EXPECT_THAT(result_or_s.status(),
              HasCompilerError("missing 1 required positional arguments 'upid'"));
}

TEST_F(PLModuleTest, upid_conversion) {
  std::string upid_str = "11285cdd-1de9-4ab1-ae6a-0ba08c8c676c";
  auto upid_or_s = md::UPID::ParseFromUUIDString(upid_str);
  ASSERT_OK(upid_or_s) << "upid should be valid.";
  auto expected_upid = upid_or_s.ConsumeValueOrDie();

  auto method_or_s = module_->GetMethod(PLModule::kUInt128ConversionId);
  ASSERT_OK(method_or_s);

  QLObjectPtr method_object = method_or_s.ConsumeValueOrDie();
  ASSERT_TRUE(method_object->type_descriptor().type() == QLObjectType::kFunction);

  auto result_or_s = std::static_pointer_cast<FuncObject>(method_object)
                         ->Call({{{"uuid_str", MakeString(upid_str)}}, {}}, ast, ast_visitor.get());
  ASSERT_OK(result_or_s);
  QLObjectPtr upid_str_object = result_or_s.ConsumeValueOrDie();
  ASSERT_TRUE(upid_str_object->type_descriptor().type() == QLObjectType::kExpr);

  std::shared_ptr<ExprObject> expr = std::static_pointer_cast<ExprObject>(upid_str_object);
  ASSERT_EQ(expr->GetExpr()->type(), IRNodeType::kUInt128);
  EXPECT_EQ(static_cast<UInt128IR*>(expr->GetExpr())->val(), expected_upid.value());
}

TEST_F(PLModuleTest, upid_conversion_fails_on_invalid_string) {
  std::string upid_str = "bad_uuid";

  auto method_or_s = module_->GetMethod(PLModule::kUInt128ConversionId);
  ASSERT_OK(method_or_s);

  QLObjectPtr method_object = method_or_s.ConsumeValueOrDie();
  ASSERT_TRUE(method_object->type_descriptor().type() == QLObjectType::kFunction);

  auto result_or_s = std::static_pointer_cast<FuncObject>(method_object)
                         ->Call({{{"uuid_str", MakeString(upid_str)}}, {}}, ast, ast_visitor.get());
  ASSERT_NOT_OK(result_or_s);
  EXPECT_THAT(result_or_s.status(), HasCompilerError(".* is not a valid UUID"));
}
}  // namespace compiler
}  // namespace carnot
}  // namespace pl
