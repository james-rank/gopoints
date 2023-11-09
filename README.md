# Points

This library is to enable pausing and crashing mid process, only for testing purposes.

*NOTE: This is a work in progress and no where near finalized and fine tuned.

## Pause Points

Sometimes you need to test a process actually gets into the desired state at a certain point, this can be difficult as subsequent methods might alter the state you're trying to check. This library gives you a way to pause the process, validate the state and then resume.

## Crash Points

Sometimes processes are able to recover themselves when a specific thread/routine has crashed, in order to easily test that behavior we can use this library to simulate crashes at specific points.