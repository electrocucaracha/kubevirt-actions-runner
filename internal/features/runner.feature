Feature: Runner resource lifecycle
  In order to run GitHub Actions workloads on KubeVirt
  As the kar runner
  I need to create, watch and delete the backing VirtualMachineInstance and DataVolume

  Background:
    Given a fresh KubeVirt runner with the default wait timeout

  Scenario Outline: Creating resources
    When resources are created with vm template "<vmTemplate>", runner name "<runnerName>" and jit config "<jitConfig>"
    Then the create result should be "<outcome>"

    Examples:
      | vmTemplate  | runnerName | jitConfig | outcome           |
      | vm-template | runnerName | jitConfig | success           |
      |             | runnerName | jitConfig | empty vm template |
      | vm-template |            | jitConfig | empty runner name |
      | vm-template | runnerName |           | empty jit config  |

  Scenario Outline: Deleting resources
    Given the app context is initialized with vmi "<vmi>" and data volume "<dataVolume>"
    When resources are deleted
    Then the operation should succeed

    Examples:
      | vmi          | dataVolume |
      | runner-xyz123 | dv-xyz123 |
      | runner-xyz123 |           |
      | runner-abc098 |           |
      | runner-xyz123 | dv-abc098 |

  Scenario: Deleting resources without an initialized app context
    Given the app context is not initialized
    When resources are deleted
    Then the operation should succeed

  Scenario Outline: Watching resources until a terminal phase
    Given a watched vmi
    When the vmi reaches phase "<lastPhase>"
    Then the watch result should be "<outcome>"

    Examples:
      | lastPhase | outcome |
      | Succeeded | success |
      | Failed    | failure |

  Scenario: Running+Ready is only a milestone before Succeeded
    Given a watched vmi that becomes Running and Ready
    When the vmi becomes Succeeded
    Then the watch result should be "success"

  Scenario: Timing out without a terminal phase
    Given a runner with a short wait timeout
    And a watched vmi that stays Running
    When the wait timeout elapses
    Then the watch should time out

  Scenario: Re-establishing the watch after the stream closes
    Given a watch that will be re-established after the first stream closes
    When the first watch stream closes after the vmi becomes Running
    And the vmi becomes Succeeded on the second watch stream
    Then the watch result should be "success"

  Scenario: Re-establishing the watch after the Running+Ready milestone
    Given a watch that will be re-established after the first stream closes
    When the vmi becomes Running and Ready then the first watch stream closes
    And the vmi becomes Succeeded on the second watch stream
    Then the watch result should be "success"

  Scenario: The wait timeout bounds a stream that keeps closing immediately
    Given a runner with a very short wait timeout
    And a watch that always returns an already closed stream for a running vmi
    When the wait timeout elapses
    Then the watch should time out

  Scenario: Exiting immediately for an already cancelled context
    Given the app context is initialized with vmi "runner-xyz123" and no data volume
    When the wait is invoked with an already cancelled context
    Then the watch should time out

  Scenario: A ready VMI reported by the initial Get is treated as a milestone
    Given a watched vmi whose initial Get already reports Running and Ready
    When the vmi becomes Succeeded
    Then the watch result should be "success"

  Scenario: Non-VMI events are ignored by the watch
    Given a watched vmi
    When an unrelated pod event is emitted
    And the vmi transitions to Succeeded
    Then the watch result should be "success"

  Scenario: Events for a different VMI name are ignored
    Given a watched vmi
    When a failed event for a different vmi name is emitted
    And the vmi transitions to Succeeded
    Then the watch result should be "success"

  Scenario: Unrecognized VMI phases are a no-op
    Given a watched vmi
    When the vmi reports an unrecognized phase
    And the vmi transitions to Succeeded
    Then the watch result should be "success"

  Scenario: Missing virtual machine template
    When resources are created referencing a nonexistent vm template
    Then the operation should fail with an error containing "failed to get KubeVirt virtual machine template"

  Scenario: VMI creation failure
    When resources are created but the VMI creation call fails
    Then the operation should fail with an error containing "failed to create runner instance"

  Scenario: Empty vm template namespace defaults to the runner namespace
    When resources are created with an empty vm template namespace
    Then the create result should be "success"

  Scenario: Creating resources when the VMI already exists
    When resources are created but the VMI already exists
    Then the create result should be "success"

  Scenario: Data volume creation failure
    When resources with a data volume template are created but the data volume creation fails
    Then the operation should fail with an error containing "cannot create data volume"

  Scenario: Creating resources that include a data volume template
    When resources with a data volume template are created successfully
    Then the operation should succeed
    And the created app context data volume name should contain "boot-disk"

  Scenario: VMI delete failure other than NotFound is only logged
    Given the app context is initialized with vmi "runner-xyz123" and no data volume
    When resources are deleted but the VMI delete call fails with a forbidden error
    Then the operation should succeed

  Scenario: Data volume delete failure other than NotFound is only logged
    Given the app context is initialized with vmi "runner-xyz123" and data volume "dv-xyz123"
    When resources are deleted but the data volume delete call fails with a forbidden error
    Then the operation should succeed

  Scenario: Watch call failure
    Given the app context is initialized with vmi "runner-xyz123" and no data volume
    When the wait is invoked but the Watch call fails
    Then the operation should fail with an error containing "failed to watch the virtual machine instance"

  Scenario: Context cancellation racing a transient Get failure
    Given the app context is initialized with vmi "runner-xyz123" and no data volume
    When the wait context is cancelled during a failing Get call
    Then the watch should time out

  Scenario: VMI disappears before the watch can be re-established
    When the vmi becomes Running then the first watch stream closes and the vmi is no longer found
    Then the operation should fail with an error containing "failed to get the virtual machine instance"
