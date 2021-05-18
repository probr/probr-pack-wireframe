@w-wel
@probes/wireframe/welcome
Feature: Protect image container registries
    This description is visible in the output,
    but serves no functional purpose

    Security Standard References:
        BCE: Basic childhood education

    Background:
        Given the config states that this test should run

    @w-wel-001
    Scenario: Ensure only family members are allowed inside
        Then "<ARRIVAL>" is "<RESPONSE>" when they come to the door

    Examples:
        | ARRIVAL        | RESPONSE |
        | mom            | welcomed |
        | dad            | welcomed |
        | sister         | welcomed |
        | brother        | welcomed |
        | thing 1        | rejected |
        | thing 2        | rejected |
        | cat in the hat | rejected |