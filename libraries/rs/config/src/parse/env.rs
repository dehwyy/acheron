use envconfig::Envconfig;

use super::Parsable;

trait Env: Parsable + Envconfig {}
impl<T: Parsable + Envconfig> Env for T {}
/// Reads environment variables and returns the parsed config.
/// If `.env` file is present, the environment variables would be extracted from it.
pub fn new_env<T>() -> T
where
    T: Env,
{
    dotenv::from_path(T::FILEPATH.to_string()).expect("Failed to load environment variables.");
    T::init_from_env().expect("Cannot initialize config from env.")
}
