import React from 'react';
import axios from 'axios';
import QRCodeImage from './QRCodeImage.js';
import RawVCard from './RawVCard.js';

// parses response data into a VCardResponse
function newVCardResponse(data) {
  if (data === undefined) {
    return null;
  }

  return {
    success: data.success,
    errors: data.errors,
    vcard_text: data.vcard_text,
    png_base64: data.png_base64
  };
}

export default class ContactQRForm extends React.Component {
    constructor(props) {
    super(props);
    this.state = {
      first: "",
      last: "",
      company_name: "",
      title: "",
      email: "",
      cell_phone: "",
      street: "",
      state: "",
      postal_code: "",
      facebook_url: "",
      twitter_handle: "",
      url: "",
      note: "",
      vcard_text: "",
      error: "",
      png_base64: "",
    };
  }

  handleChange = e => {
    this.setState({ [e.target.name]: e.target.value });
  };

  handleResponseData = (data) => {
    const vCardResponse = newVCardResponse(data);
    if (vCardResponse.success) {
      this.setState({vcard_text: vCardResponse.vcard_text});
      this.setState({png_base64: vCardResponse.png_base64});
    } else {
      this.setState({error: vCardResponse.errors});
    }
  };

  onSubmit = (e) => {
    e.preventDefault();
    this.setState({vcard_text: ""});
    this.setState({error: ""});
    this.setState({png_base64: ""});

    // create a VCardRequest request object from the form
    const vCardRequest = {
      first: this.state.first,
      last: this.state.last,
      company_name: this.state.p,
      title: this.state.title,
      email: this.state.email,
      cell_phone: this.state.cell_phone,
      street: this.state.street,
      state: this.state.state,
      postal_code: this.state.postal_code,
      facebook_url: this.state.facebook_url,
      twitter_handle: this.state.twitter_handle,
      url: this.state.url,
      note: this.state.note
    };

    const handleResponseData = this.handleResponseData;
    axios.post('/api/v1/vcard/create', vCardRequest)
      .then(response => {
        handleResponseData(response.data);
      }).catch(error => {
        handleResponseData(error.response.data);
      });
  };

  render() {
    const { first, last, company_name, title, email, cell_phone,
      street, state, postal_code, facebook_url, twitter_handle,
      url, note, vcard_text, error, png_base64 } = this.state;

      // show error message
    let errorMsg = "";
    if (error.length) {
      errorMsg =
      <span>
        <br/>{error}
      </span>;
    }

    // show the QR Code image
    let qrCode = "";
    if (png_base64.length > 0) {
      qrCode = <QRCodeImage png_base64={png_base64} />
    }

    // show the raw vCard text
    let vCardRawText = "";
    if (vcard_text.length > 0) {
      vCardRawText = <RawVCard vcard_text={vcard_text} />
    }

    return (
      <div>
         <form name="contactQRForm" onSubmit={this.onSubmit}>
          <p>
            <label>
              First Name:<br />
              <input type="text" name="first" value={first} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Last Name:<br />
              <input type="text" name="last" value={last} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Company Name:<br />
              <input type="text" name="company_name" value={company_name} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Title:<br />
              <input type="text" name="title" value={title} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Email:<br />
              <input type="text" name="email" value={email} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Cell Phone:<br />
              <input type="text" name="cell_phone" value={cell_phone} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Address:<br />
              TODO
            </label>
          </p>

          <p>
            <label>
              Facebook Profile URL:<br />
              <input type="text" name="facebook_url" value={facebook_url} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Twitter Handle:<br />
              <input type="text" name="twitter_handle" value={twitter_handle} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Website URL:<br />
              <input type="text" name="url" value={url} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <label>
              Note:<br />
              <textarea name="note" value={note} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <button type="submit">Create QR Code</button>
            {errorMsg}
          </p>
        </form>
        {qrCode}
        {vCardRawText}
      </div>
    );
  }
}