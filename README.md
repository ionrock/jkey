# jkey

```
$ jkey --help
jkey

Takes stdin and a key and prints out all instances of that key.

Usage: jkey
  -k string
        The json key to search for
$ cat tmp.json 
{
    "baz": "bar",
    "a": {
        "b": {
            "foo": "bar"
        }
    },
    "b": {
        "c": {
            "d": {
                "foo": "bar"
            }
        }
    }
}
$ cat tmp.json | jkey -k foo
Found the key at .a.b.foo
Found the key at .b.c.d.foo
Done searching
```
