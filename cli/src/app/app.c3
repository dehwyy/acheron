module app;
import std::io;
import command;
import gql;


<*
	@param [in] args `Command line arguments.`
	@require args.len > 1, `Command not provided.`
*>
fn void run(String[] args) {
	command::Command! command = command::command_from_string(args[1]);

	if (catch excuse = command) {
		case CommandCastError.COMMAND_NOT_FOUND:
		 	command::print_command_not_found();
			command::print_help_command();

		return;
	}

	switch (command!!) {
		case Command.HELP:
			command::print_help_command();
		case Command.GQL:
			if (catch e = command::execute_command(gql::command())) {
				io::printn(e);
			}
	}

}
