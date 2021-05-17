@w-fra
@probes/wireframe/wireframe
Feature: Protect image container registries
    This description is visible in the output,
    but serves no functional purpose

    Security Standard References:
        BCE: Basic childhood education

    Background:
        Given the config states that this test should run

    @k-cra-003
    Scenario: Ensure only family members are allowed inside
        Then "<VISITOR>" is "<RESPONSE>" when they come to the door

    Examples:
        | VISITOR        | RESPONSE |
        | mom            | welcomed |
        | dad            | welcomed |
        | sister         | welcomed |
        | brother        | welcomed |
        | thing 1        | rejected |
        | thing 2        | rejected |
        | cat in the hat | rejected |