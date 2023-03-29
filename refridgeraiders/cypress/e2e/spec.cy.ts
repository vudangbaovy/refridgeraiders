describe('Login Test', () => {
  it('Visits the Kitchen Sink', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.contains('First time at KitchenRescue?').click()
    cy.url().should('include', '/register')
  })
})