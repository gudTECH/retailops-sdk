package verify

import (
  "os"

  "fmt"
  "strings"

  p "path"
  fp "path/filepath"

)

func doLocalVerify(cliExec CLIExecution) (err error) {
  var examples []SchemaExample

  if cliExec.SchemaPathIsDir {
    examples,err = allExamples(cliExec)
  } else {
    examples,err = examplesForSchema(cliExec.SchemaPath, cliExec)
  }

  if err != nil {
    return
  } else if len(examples) == 0 {
    err = fmt.Errorf("no test JSON files found. try `--help` for more information")
    return
  }

  fmt.Println(len(examples),"TESTS TO BE GENERATED")
  errorCount := 0
  for index,example := range examples {
    err = doVerify(cliExec, index, example)
    if err != nil {
        errorCount ++
    }
  }

  if errorCount > 0 {
      err = fmt.Errorf("\nLocal verify experienced %d errors", errorCount)
      return
  }

  return
}

func doVerify(cliExec CLIExecution, index int, testCase SchemaExample) (err error) {
  var thereWasAnError bool
  fmt.Println(HR)
  fmt.Printf("TEST %d (%s)", index+1, p.Base(testCase.ExamplePath))

  err = loadFilesAndMakeRequest(cliExec.BaseURL, testCase.SchemaPath, testCase.ExamplePath, cliExec.Verbose)
  if err != nil {
    fmt.Printf("\rTEST %d (%s) FAILED: %s\n", index+1, p.Base(testCase.ExamplePath), err.Error())
    if cliExec.StopOnError {
      os.Exit(1)
    } else {
      thereWasAnError = true
    }
  } else {
    if cliExec.Verbose {
      fmt.Println("")
    }
    fmt.Printf("\rTEST %d (%s) WAS A SUCCESS\n", index+1, p.Base(testCase.ExamplePath))

    if cliExec.Verbose {
      fmt.Println("")
    }
  }

  if thereWasAnError && cliExec.StopOnError {
    err = fmt.Errorf("AT LEAST ONE OF THE TEST CASES FAILED")
    return
  }

  return
}

func loadFilesAndMakeRequest(baseUrl, schemaPath, examplePath string, verbose bool) (err error) {

  f,err := os.Open(schemaPath)
  if err != nil {
    return
  }

  exampleF,err := os.Open(examplePath)
  if err != nil {
    return
  }

  err = Request(baseUrl, "RETAILOPS_SDK", f, exampleF, verbose, 200)
  if err != nil {
    return
  }

  _,err = exampleF.Seek(0,0)
  if err != nil {
    return
  }
  _,err = f.Seek(0,0)
  if err != nil {
    return
  }
  // send invalid test key, should receive 401 error to pass
  err = Request(baseUrl, "KDS_SPOLIATER", f, exampleF, verbose, 401)
  if err == nil {
    return err
  }

  return
}

func examplesForSchema(schemaPath string, cliExec CLIExecution) (verifications []SchemaExample, err error) {
  dirname,filename := fp.Split(schemaPath)
  exampleFilename := strings.Replace(filename, ".json", "", -1)
  exampleFilenameGlob := fmt.Sprintf("%s_ex_*.json", exampleFilename)

  if cliExec.ExamplesPathIsDir {
      dirname = cliExec.ExamplesPath
  }
  pathGlob := p.Join(dirname, exampleFilenameGlob)

  examplePaths,err := fp.Glob(pathGlob)
  if err != nil {
    return
  }
  verifications = make([]SchemaExample,0)
  for _,exPath := range examplePaths {
    verifications = append(verifications, SchemaExample{
      SchemaPath: schemaPath,
      ExamplePath: exPath,
    })
  }

  return
}

func allExamples(cliExec CLIExecution) (verifications []SchemaExample, err error) {
  verifications = make([]SchemaExample,0)
  dirname := cliExec.SchemaPath
  filter := cliExec.SchemaFilter
  var allSchemasGlob string
  var allSchemaPaths []string

    if filter == "" {
        allSchemasGlob = p.Join(dirname, "*v1.json")
        allSchemaPaths, err = fp.Glob(allSchemasGlob)
        if err != nil {
            return
        }
    } else {
        for _, action := range cliExec.CertifyActions {
            globPath := p.Join(dirname, fmt.Sprintf("*%s*v1.json", action))
            path, err := fp.Glob(globPath)
            if err != nil {
                return nil, fmt.Errorf("No schemas found for pattern: `%s`", globPath)
            }
            allSchemaPaths = append(allSchemaPaths, path...)
        }
    }
    
    if len(allSchemaPaths) == 0 {
        err = fmt.Errorf("`%s` did not contain schemas", dirname)
        return
    }

  for _,schemaPath := range allSchemaPaths {
    exs,err := examplesForSchema(schemaPath, cliExec)
    if err != nil {
      return nil,err
    }

    for _,ex := range exs {
      verifications = append(verifications, ex)
    }
  }

  return
}
