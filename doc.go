// Package lameduck is a simple API that provides a very basic framework for
// setting up lame duck mode for a service. The basic idea is that you create
// a LameDuckHandler by giving it the function you want to call when your
// service enters lame duck mode. You then tell that LameDuckHandler how it
// it should be notified to enter lameduck mode, and then you start it.
// Starting a LameDuckHandler spawns a go routine per handler listener.
// A basic example is as follows:
//
//  // lameDuckFn = our function to call when entering lame duck handler
//  lameduck.NewLameDuckHandler(
//    lameduck.EnterLameDuckMode(func(){
//      fmt.Println("entering lameduck mode!")
//    }),
//  ).WithSigINTHandler().Go()
package lameduck
