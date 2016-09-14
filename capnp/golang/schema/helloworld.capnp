using Go = import "/go.capnp";
@0xbbd09ee5a468f60c;

$Go.package("greeter");
$Go.import("foo/greeter");

# The greeting service definition.
interface Greeter {
  # Sends a greeting
  sayHello @0 (name :Text) -> (rep :Text);
}
