
# testing 

writing a test for a function

important steps:
- Test needs to be in the file with the name ``xxx_test.go``
  (important as the test won't be recognized if it's in a different file)
  
- import the testing package

```go
import "testing"
```

- create a test function that starts with `Test`
- call the function with `t.Run()`
- Test function takes one argument of type `*testing.T`
```go
func TestAdd(t *testing.T) {
}
```

## Sub-tests
can run sub-tests with `t.Run("sub-test", func(t *testing.T) {})` inside a test function.
```go
func TestAdd(t *testing.T) {
	t.Run("sub-test", func(t *testing.T) {
		
	})
}
```

- Example of Hello subtests for different conditions
```go
import "testing"

func TestIf(t *testing.T) {
	// subtests
	t.Run("saying hello to ppl", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello! Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	// subtests
	t.Run("Say \"Hello! World\" when empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello! World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("check language specific functionality", func(t *testing.T) {
		got := Hello("Asta", "Spanish")
		want := "Hola! Asta"
		assertCorrectMessage(t, got , want)
	})

}

```

## Discipline
Let's go over the cycle again

1. Write a test

2. Make the compiler pass

3. Run the test, see that it fails and check the error message is meaningful

4. Write enough code to make the test pass

5. Refactor
