module command;
import libc;
import std::io;
import gql;

enum Command: char {
	HELP,
	GQL
}

fault CommandCastError {
	COMMAND_NOT_FOUND
}
<*
	@param [in] command `Potential command provided via terminal.`
	@return `Command ensured that CLI supports it.`
	@return! CommandCastError.COMMAND_NOT_FOUND `If command doesn't match any supported command.`
*>
fn Command! command_from_string(String command) @public {
	switch (command) {
		case "help":
			return Command.HELP;
		case "gql":
			return Command.GQL;
		default:
			return CommandCastError.COMMAND_NOT_FOUND?;
	}
}

// Prints help command
fn void print_help_command() @public {
	io::printfn(
`Usage: cli.exe <command> [<args>]

  Commands:
	help - prints help command
	gql - %s
`,
	gql::description(),
	);
}

//
fn void print_command_not_found() @public {
	io::printn("Command not found");
}

fault CommandFail {
	EXECUTION_FAILED
}

<*
	@param [in] string_to_execute `Command represented as String. Would be casted to ZString.`
	@return! CommandFail.EXECUTION_FAILED `If 'system' fn returns status == 1.`
*>
fn void! execute_command(String string_to_execute) @public {
	int status_code = libc::system(String.zstr_copy(string_to_execute));

	// I wonder whether it should be more like `status_code != 0`?
	if (status_code == 1) {
		return CommandFail.EXECUTION_FAILED?;
	}
}
