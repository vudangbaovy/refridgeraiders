describe('Login Test', () => {
  it('link navigation', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.contains('First time at KitchenRescue?').click()
    cy.url().should('include', '/register')
  })
})

describe('Profile Component', () => {
  it('should retrieve user profile on init', () => {
    cy.intercept('POST', 'http://localhost:3000/user', (req) => {
      req.reply({ user: 'vyvooz' });
    }).as('getUser');

    cy.visit('http://localhost:4200/profile');

    cy.wait('@getUser').then((interception) => {
      expect(interception.response.body).to.have.property('user', 'vyvooz');
    });
  });
});
