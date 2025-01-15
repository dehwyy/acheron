module cli;
import app;

fn int main(String[] args)
{
	App app = app::new_app();
	app.run(args)!!;

	return 0;
}
