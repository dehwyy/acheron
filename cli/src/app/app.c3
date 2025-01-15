module app;
import std::io;

fn App new_app() {
	App app;
	app.dev  = true;
	return app;
}

struct App {
	bool dev;
}

fn void! App.run(&self, String[] args) {
	foreach (index, &arg : args[1..]) {
		io::printfn("arg: %s", *arg);
	}
}
