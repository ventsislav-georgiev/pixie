<?php

/**
 * This file is automatically generated. Use 'arc liberate' to rebuild it.
 *
 * @generated
 * @phutil-library-version 2
 */
phutil_register_library_map(array(
  '__library_version__' => 2,
  'class' => array(
    'ArcaistTypescriptLinter' => 'ts/lint/ArcaistTypescriptLinter.php',
    'ArcanistBaseGenCheckerTestEngine' => 'pixielabs/unit/ArcanistBaseGenCheckerTestEngine.php',
    'ArcanistClangFormatLinter' => 'clang_format/lint/ArcanistClangFormatLinter.php',
    'ArcanistESLintLinter' => 'js/lint/ArcanistESLintLinter.php',
    'ArcanistGoGenCheckerTestEngine' => 'pixielabs/unit/ArcanistGoGenCheckerTestEngine.php',
    'ArcanistGoImportsLinter' => 'pixielabs/lint/ArcanistGoImportsLinter.php',
    'ArcanistGoVetLinter' => 'golang/lint/ArcanistGoVetLinter.php',
    'ArcanistGolangCiLinter' => 'pixielabs/lint/ArcanistGolangCiLinter.php',
    'ArcanistGraphqlGenCheckerTestEngine' => 'pixielabs/unit/ArcanistGraphqlGenCheckerTestEngine.php',
    'ArcanistProtoBreakCheckLinter' => 'pixielabs/lint/ArcanistProtoBreakCheckLinter.php',
    'ArcanistProtoGenCheckerTestEngine' => 'pixielabs/unit/ArcanistProtoGenCheckerTestEngine.php',
    'ArcanistShellCheckLinter' => 'shellcheck/lint/ArcanistShellCheckLinter.php',
    'ArcanistShellCheckLinterTestCase' => 'shellcheck/lint/__tests__/ArcanistShellCheckLinterTestCase.php',
  ),
  'function' => array(),
  'xmap' => array(
    'ArcaistTypescriptLinter' => 'ArcanistExternalLinter',
    'ArcanistBaseGenCheckerTestEngine' => 'ArcanistUnitTestEngine',
    'ArcanistClangFormatLinter' => 'ArcanistExternalLinter',
    'ArcanistESLintLinter' => 'ArcanistExternalLinter',
    'ArcanistGoGenCheckerTestEngine' => 'ArcanistBaseGenCheckerTestEngine',
    'ArcanistGoImportsLinter' => 'ArcanistExternalLinter',
    'ArcanistGoVetLinter' => 'ArcanistExternalLinter',
    'ArcanistGolangCiLinter' => 'ArcanistExternalLinter',
    'ArcanistGraphqlGenCheckerTestEngine' => 'ArcanistBaseGenCheckerTestEngine',
    'ArcanistProtoBreakCheckLinter' => 'ArcanistExternalLinter',
    'ArcanistProtoGenCheckerTestEngine' => 'ArcanistBaseGenCheckerTestEngine',
    'ArcanistShellCheckLinter' => 'ArcanistExternalLinter',
    'ArcanistShellCheckLinterTestCase' => 'PhutilTestCase',
  ),
));