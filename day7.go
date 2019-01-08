package main

// day 7 outline

// struct AssemblyStep
// - ID
// - after: Array of required Assembly steps

// algo outline
// - look for
// - parse each line into Map[ID] -> req steps. Append etc
//
// find first step: findCandidates([]).
//
//

// findCandidates(executionOrder)
// - find steps that haven't been executed:
// - loop through steps
// -- if already executed, skip
// -- if prereqs not in executionOrder, skip
// -- pick earliest in alphabet
//
