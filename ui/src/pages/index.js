import React from 'react'

import Layout from '../components/layout'
import ContactQRForm from '../components/contactQRForm'

const IndexPage = () => (
  <Layout>
    <p>Show someone a QR Code with your contact information in vCard format instead of sending someone a SMS (<em>because that's just too much work</em>).
    More info at <a href="https://github.com/aphexddb/contactqr">https://github.com/aphexddb/contactqr</a>.</p>

    <ContactQRForm/>

  </Layout>
)

export default IndexPage
