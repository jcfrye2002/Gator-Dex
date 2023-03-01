import { LoginComponent } from "src/app/login/login.component"

describe('LoginComponent', () => {
  it('mounts', () => {
    cy.mount(LoginComponent)
  })
})

it('when the button is pressed, form is submitted', () => {
  cy.mount(LoginComponent)
  cy.get('.button')
})