import { createOutputSpy } from "cypress/angular"
import { LoginComponent } from "src/app/login/login.component"

describe('LoginComponent', () => {
  it('mounts', () => {
    cy.mount(LoginComponent)
  })
})

it('Form is fillable and button is clikable', () => {
  cy.mount(LoginComponent)
  cy.get('[data-cy=emailBox]').type("user@email.com")
  cy.get('[data-cy=passwordBox]').type("password")
  cy.get('[data-cy=submitButton]').click({force:true})
})