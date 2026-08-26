Feature: Root command initialization
  The kar root command creates, waits for and deletes runner resources

  Scenario Outline: initialization process
    Given a mock runner
    When the root command is executed with flag "<flag>" value "<value>" and induced failure "<failure>"
    Then the command execution should "<outcome>"

    Examples:
      | flag | value       | failure | outcome |
      |      |             | none    | succeed |
      | -c   | test config | none    | succeed |
      | -t   | vm template | none    | succeed |
      | -n   | kubevirt    | none    | succeed |
      | -r   | runner name | none    | succeed |
      |      |             | create  | fail    |
      |      |             | delete  | fail    |
      |      |             | wait    | fail    |
