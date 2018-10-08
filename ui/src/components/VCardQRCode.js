import React from 'react'
import './VCardQRCode.css';

export default class VCardQRCode extends React.Component {

  render() {
    const componentClasses = ['vcard-qrcode-component'];
    const { vcard_text, png_base64, show } = this.props;

    if (show) { componentClasses.push('show'); }

    return (
      <div className={componentClasses.join(' ')}>
        <div className="qrcode-wrapper">
          <div>
            <img src={png_base64} alt="QR Code" />
          </div>
          <div class="instructions">
            Tap and hold the QR Code to save it to your device.
          </div>
        </div>
        <div class="raw-data">
          <p>vCard data in <a href="https://tools.ietf.org/html/rfc6350">RFC 6350</a> format:</p>
          <pre>{vcard_text}</pre>
        </div>
      </div>
    );
  }
}