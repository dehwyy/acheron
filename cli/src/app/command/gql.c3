module gql;
import libc;

fn String command() {
	return "cd apps/stream_broadcaster && go run github.com/99designs/gqlgen generate";
}

fn String description() {
	return "generates .go files from .gql for 'apps/stream_broadcaster' microservice";
}
