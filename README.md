    NAME

    sm-mustache - Mustache template command line

    SYNOPSIS

      sm-mustache [options]

    OPTIONS

      -data="{key1=value1}:~{key2=value2}": key=value data pair, overrides json file data
      -json="{string}": json data string
      -output="{file}": path to output file
      -template="{file}": path to template file

      -data|--data "{key1=value1:~key2=value2}

        Assign a value to a given key for the template rendering

      -json|--json "{string}"

        To be implemented

      -output|--output "{file}"

        specify the output file to write the rendered template to

      -template|--template "{file}"

        specify the template file to use

    DESCRIPTION

    sm-mustache will read the given template file, and using the given data render
    it to the specified output file location.

