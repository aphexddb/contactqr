module.exports = {
  siteMetadata: {
    title: 'Contact QR',
  },
  plugins: [
    'gatsby-plugin-react-helmet',
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        name: 'gatsby-starter-default',
        short_name: 'starter',
        start_url: '/',
        background_color: '#38a5fe',
        theme_color: '#38a5fe',
        display: 'minimal-ui',
        icon: 'src/images/favicon-contact.png', // This path is relative to the root of the site.
      },
    },
    'gatsby-plugin-offline',
  ],
}
