import React from 'react'

export default class RawVCard extends React.Component {

  render() {
    const { vcard_text } = this.props;

    return (
      <div>
        <p>vCard data in <a href="https://tools.ietf.org/html/rfc6350">RFC 6350</a> format:</p>
        <pre>{vcard_text}</pre>
      </div>
    );
  }
}