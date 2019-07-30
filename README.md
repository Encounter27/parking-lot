step 1: system should have go installed.

step2: install direnv
    https://tech.willandskill.se/install-direnv-on-ubuntu-16-04/

    OR set the GOPATH manually from root directory

    [parking_lot]$ export GOPATH=$(pwd)

step3: run the bin/setup
    [parking_lot]$ ./bin/setup

    # you can see the bellow output:

    github.com/google/uuid (download)
    ?   	parking-lot-app	[no test files]
    ok  	parking-lot-app/command	0.009s	coverage: 88.7% of statements
    ok  	parking-lot-app/parking	0.007s	coverage: 60.1% of statements
    ok  	parking-lot-app/utils	0.016s	coverage: 100.0% of statements

step 4: To run the specs :
    [parking_lot]$ bin/run_functional_tests

    You can see the below output:
    
    Finished in 17.64 seconds (files took 0.14673 seconds to load)
    6 examples, 0 failures

step 5: [Optional]
    [parking_lot]$ ./dev_setup.sh

    It will be helpful to explore the code using vscode IDE, as this will enable the intelligence of vscode for Go.

[please note] I have changed a bit the fuctional spec command, to read from file.

 pty = PTY.spawn("parking_lot < #{File.join(File.dirname(__FILE__), '..', 'fixtures', 'file_input.txt')}")
                              ^
                              |

                        Without this redirect operator the executable will not be able to read from file.

Also please use the make file at the root directory to build, run and run-unit tests.
