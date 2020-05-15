#include "src/stirling/pgsql/parse.h"

#include <string>
#include <utility>

#include "src/common/testing/testing.h"
#include "src/stirling/common/stitcher.h"
#include "src/stirling/pgsql/types.h"

namespace pl {
namespace stirling {
namespace pgsql {

using ::testing::AllOf;
using ::testing::ElementsAre;
using ::testing::Field;
using ::testing::IsEmpty;
using ::testing::StrEq;

template <typename TagMatcher, typename LengthType, typename PayloadMatcher>
auto IsRegularMessage(TagMatcher tag, LengthType len, PayloadMatcher payload) {
  return AllOf(Field(&RegularMessage::tag, tag), Field(&RegularMessage::len, len),
               Field(&RegularMessage::payload, payload));
}

const std::string_view kQueryTestData =
    CreateStringView<char>("Q\000\000\000\033select * from account;\000");
const std::string_view kStartupMsgTestData = CreateStringView<char>(
    "\x00\x00\x00\x54\x00\003\x00\x00user\x00postgres\x00"
    "database\x00postgres\x00"
    "application_name\x00psql\x00"
    "client_encoding\x00UTF8\x00\x00");

TEST(PGSQLParseTest, BasicMessage) {
  std::string_view data = kQueryTestData;
  RegularMessage msg = {};
  EXPECT_EQ(ParseState::kSuccess, ParseRegularMessage(&data, &msg));
  EXPECT_THAT(
      msg, IsRegularMessage(Tag::kQuery, 27, CreateStringView<char>("select * from account;\0")));
}

auto IsNV(std::string_view name, std::string_view value) {
  return AllOf(Field(&NV::name, name), Field(&NV::value, value));
}

TEST(PGSQLParseTest, StartupMessage) {
  std::string_view data = kStartupMsgTestData;
  StartupMessage msg = {};
  EXPECT_EQ(ParseState::kSuccess, ParseStartupMessage(&data, &msg));
  EXPECT_EQ(84, msg.len);
  EXPECT_EQ(3, msg.proto_ver.major);
  EXPECT_EQ(0, msg.proto_ver.minor);
  EXPECT_THAT(msg.nvs,
              ElementsAre(IsNV("user", "postgres"), IsNV("database", "postgres"),
                          IsNV("application_name", "psql"), IsNV("client_encoding", "UTF8")));
  EXPECT_THAT(data, IsEmpty());
}

const std::string_view kRowDescTestData = CreateStringView<char>(
    "T\000\000\000\246"
    "\000\006"
    "Name"
    "\000\000\000\004\356\000\002\000\000\000\023\000@\377\377\377\377\000\000"
    "Owner"
    "\000\000\000\000\000\000\000\000\000\000\023\000@\377\377\377\377\000\000"
    "Encoding"
    "\000\000\000\000\000\000\000\000\000\000\023\000@\377\377\377\377\000\000"
    "Collate"
    "\000\000\000\004\356\000\005\000\000\000\023\000@\377\377\377\377\000\000"
    "Ctype"
    "\000\000\000\004\356\000\006\000\000\000\023\000@\377\377\377\377\000\000"
    "Access "
    "privileges\000\000\000\000\000\000\000\000\000\000\031\377\377\377\377\377\377\000\000");

TEST(PGSQLParseTest, RowDesc) {
  std::string_view data = kRowDescTestData;
  RegularMessage msg = {};
  EXPECT_EQ(ParseState::kSuccess, ParseRegularMessage(&data, &msg));
  EXPECT_EQ(Tag::kRowDesc, msg.tag);
  EXPECT_EQ(166, msg.len);
  EXPECT_THAT(ParseRowDesc(msg.payload),
              ElementsAre("Name", "Owner", "Encoding", "Collate", "Ctype", "Access privileges"));
}

const std::string_view kDataRowTestData = CreateStringView<char>(
    "D"
    "\000\000\000F"
    "\000\006"
    "\000\000\000\010postgres"
    "\000\000\000\010postgres"
    "\000\000\000\004UTF8"
    "\000\000\000\nen_US.utf8"
    "\000\000\000\nen_US.utf8"
    "\377\377\377\377");

TEST(PGSQLParseTest, DataRow) {
  std::string_view data = kDataRowTestData;
  RegularMessage msg = {};
  EXPECT_EQ(ParseState::kSuccess, ParseRegularMessage(&data, &msg));
  EXPECT_EQ(Tag::kDataRow, msg.tag);
  EXPECT_EQ(70, msg.len);
  EXPECT_THAT(ParseDataRow(msg.payload), ElementsAre("postgres", "postgres", "UTF8", "en_US.utf8",
                                                     "en_US.utf8", std::nullopt));
}

auto ParamIs(FmtCode fmt_code, std::string_view value) {
  return AllOf(Field(&Param::format_code, fmt_code), Field(&Param::value, value));
}

TEST(PGSQLParseTest, ParseBindRequest) {
  std::string_view bind_msg_payload = CreateStringView<char>(
      // destination portal name
      "destination portal\x00"
      // source prepared statement name
      "source prepared statement\x00"
      // Parameter format code count
      "\x00\x01"
      // 1st parameter format code
      "\x00\x01"
      // Parameter count
      "\x00\x02"
      // 1st parameter value length
      "\x00\x00\x00\x05"
      // 1st parameter value
      "\x4a\x61\x73\x6f\x6e"
      // 2nd parameter value length
      "\x00\x00\x00\x03"
      // 2nd parameter value
      "\xaa\xbb\xcc"
      // Result column format code count
      "\x00\x02"
      // 1st result column format code
      "\x00\x00"
      // 2nd result column format code
      "\x00\x01");
  BindRequest bind_req;
  EXPECT_EQ(ParseState::kSuccess, ParseBindRequest(bind_msg_payload, &bind_req));
  EXPECT_THAT(bind_req.dest_portal_name, StrEq("destination portal"));
  EXPECT_THAT(bind_req.src_prepared_stat_name, StrEq("source prepared statement"));
  EXPECT_THAT(bind_req.params, ElementsAre(ParamIs(FmtCode::kBinary, "\x4a\x61\x73\x6f\x6e"),
                                           ParamIs(FmtCode::kBinary, "\xaa\xbb\xcc")));
  EXPECT_THAT(bind_req.res_col_fmt_codes, ElementsAre(FmtCode::kText, FmtCode::kBinary));
}

TEST(PGSQLParseTest, ParseParamDesc) {
  std::string_view payload = CreateStringView<char>(
      // Parameter count
      "\x00\x03"
      // 1st type oid
      "\x00\x00\x00\x19"
      // 2nd type oid
      "\x00\x00\x00\x19"
      // 2rd type oid
      "\x00\x00\x00\x19");
  ParamDesc param_desc;
  EXPECT_EQ(ParseState::kSuccess, ParseParamDesc(payload, &param_desc));
  EXPECT_THAT(param_desc.type_oids, ElementsAre(25, 25, 25));
}

TEST(PGSQLParseTest, ParseParse) {
  std::string_view payload = CreateStringView<char>(
      "test\x00"
      "SELECT * FROM person WHERE first_name=$1\x00"
      // Parameter type OIDs count
      "\x00\x01"
      "\x00\x00\x00\x19");
  Parse parse;
  EXPECT_EQ(ParseState::kSuccess, ParseParse(payload, &parse));
  EXPECT_THAT(parse.stmt_name, StrEq("test"));
  EXPECT_THAT(parse.query, StrEq("SELECT * FROM person WHERE first_name=$1"));
  EXPECT_THAT(parse.param_type_oids, ElementsAre(25));
}

const std::string_view kDropTableCmplMsg =
    CreateStringView<char>("C\000\000\000\017DROP TABLE\000");

TEST(FindFrameBoundaryTest, FindTag) {
  EXPECT_EQ(0, FindFrameBoundary(kDropTableCmplMsg, 0));
  EXPECT_EQ(0, FindFrameBoundary(kRowDescTestData, 0));
  EXPECT_EQ(0, FindFrameBoundary(kDataRowTestData, 0));
  const std::string data = absl::StrCat("aaaaa", kDataRowTestData);
  EXPECT_EQ(5, FindFrameBoundary(data, 0));
}

}  // namespace pgsql
}  // namespace stirling
}  // namespace pl
