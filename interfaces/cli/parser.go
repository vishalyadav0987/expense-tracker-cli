package cli

import "errors"

type Command struct {
	Name string
	Args map[string]string
}

func Parse(args []string) (*Command, error) {

	if len(args) < 2 {
		return nil, errors.New("no command provided")
	}
	cmd := &Command{
		Name: args[1],
		Args: make(map[string]string),
	}

	for i := 2; i < len(args); i++ {

		arg := args[i]

		if len(arg) > 2 && arg[:2] == "--" {
			key := arg[2:]

			// check if value exists
			if i+1 < len(args) {
				cmd.Args[key] = args[i+1]
				i++ // skip value
			} else {
				return nil, errors.New("flag value missing for " + key)
			}
		}
	}

	return cmd, nil
}
