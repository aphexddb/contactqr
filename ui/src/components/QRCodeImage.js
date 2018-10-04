import React from 'react'

export default class QRCodeImage extends React.Component {

  render() {
    const { png_base64 } = this.props;

    return (
      <div>
        <img src={png_base64} alt="QR Code" />
      </div>
    );
  }
}