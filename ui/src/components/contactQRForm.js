import React from 'react';
import axios from 'axios';
import './formLayout.css';
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
      city: "",
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
      city: this.state.city,
      state: this.state.state,
      postal_code: this.state.postal_code,
      facebook_url: this.state.facebook_url,
      twitter_handle: this.state.twitter_handle,
      url: this.state.url,
      note: this.state.note
    };

    const handleResponseData = this.handleResponseData;
    axios.post('http://localhost:8080/api/v1/vcard/create', vCardRequest)
      .then(response => {
        handleResponseData(response.data);
      }).catch(error => {
        console.log("error:", error);
        handleResponseData(error.response.data);
      });
  };

  render() {
    const { first, last, company_name, title, email, cell_phone,
      street, city, state, postal_code, facebook_url, twitter_handle,
      url, note, vcard_text, error, png_base64 } = this.state;

      // show error message
    let errorMsg = "";
    if (error.length) {
      errorMsg =
      <div className="alert alert-warning" role="alert">
        {error}
      </div>;
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
      <div id="container">
        <form name="contactQRForm" onSubmit={this.onSubmit} className="needs-validation" noValidate>

          <div className="form-row">
            <div className="col-md-3 mb-3">
              <label htmlFor="first">First name</label>
              <input type="text" className="form-control" placeholder="David" name="first" value={first} onChange={this.handleChange} />
            </div>
            <div className="col-md-3 mb-3">
              <label htmlFor="last">Last name</label>
              <input type="text" className="form-control" placeholder="Bowie" name="last" value={last} onChange={this.handleChange} />
            </div>
            <div className="col-md-4 mb-3">
              <label htmlFor="email">Email</label>
              <input type="text" className="form-control" placeholder="dbowie@anothercastle.com" name="email" value={email} onChange={this.handleChange} />
            </div>
            <div className="col-md-2 mb-3">
              <label htmlFor="cell_phone">Cell Phone</label>
              <input type="text" className="form-control" placeholder="314-222-1234" name="cell_phone" value={cell_phone} onChange={this.handleChange} />
            </div>
          </div>

          <div className="form-row">
            <div className="col-md-3 mb-3">
              <label htmlFor="company_name">Company name</label>
              <input type="text" className="form-control" placeholder="Another Castle Games" name="company_name" value={company_name} onChange={this.handleChange} />
            </div>
            <div className="col-md-3 mb-3">
              <label htmlFor="title">Title</label>
              <input type="text" className="form-control" placeholder="Mushroom Farmer" name="title" value={title} onChange={this.handleChange} />
            </div>
            <div className="col-md-2 mb-3">
              <label htmlFor="twitter_handle">Twitter</label>
              <div className="input-group">
                <div className="input-group-prepend">
                  <span className="input-group-text" id="inputGroupPrepend">@</span>
                </div>
                <input type="text" className="form-control" placeholder="star" name="twitter_handle" value={twitter_handle} onChange={this.handleChange} />
              </div>
            </div>
            <div className="col-md-4 mb-3">
              <label htmlFor="facebook_url">Facebook URL</label>
              <input type="text" className="form-control" placeholder="https://www.facebook.com/dbowie" name="facebook_url" value={facebook_url} onChange={this.handleChange} />
            </div>
          </div>

          <div className="form-row">
            <div className="col-md-12 mb-12">
              <label>Address</label>
            </div>
            <div className="col-md-4 mb-3">
              <input type="text" className="form-control" placeholder="350 Fifth Avenue" name="street" value={street} onChange={this.handleChange} />
            </div>
            <div className="col-md-4 mb-3">
              <input type="text" className="form-control" placeholder="New York" name="city" value={city} onChange={this.handleChange} />
            </div>
            <div className="col-md-2 mb-3">
              <input type="text" className="form-control" placeholder="NY" name="state" value={state} onChange={this.handleChange} />
            </div>
            <div className="col-md-2 mb-3">
              <input type="text" className="form-control" placeholder="10118" name="postal_code" value={postal_code} onChange={this.handleChange} />
            </div>
          </div>

          <div className="form-row">
            <div className="col-md-4 mb-3">
              <label htmlFor="url">Website URL</label>
              <input type="text" className="form-control" placeholder="https://www.anothercastle.com" name="url" value={url} onChange={this.handleChange} />
            </div>
            <div className="col-md-8 mb-3">
              <label htmlFor="note">Note</label>
              <input type="text" className="form-control" placeholder="" name="note" value={note} onChange={this.handleChange} />
            </div>
          </div>

          <button className="btn btn-primary" type="submit">Create QR Code</button>
          {errorMsg}
        </form>

        {qrCode}
        {vCardRawText}
      </div>
    );
  }
}