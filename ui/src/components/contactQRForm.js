import React from 'react';
import axios from 'axios';
import './ContactQRForm.css';
import VCardQRCode from './VCardQRCode';
import ErrorMessage from './ErrorMessage';
import CreateAnotherButton from './CreateAnotherButton';

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
    this.state = this.getDefaultState();
  };

  getDefaultState = () => {
    return {
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
      hide: false
    };
  };

  resetFormState = () => {
    this.setState(this.getDefaultState());
  };

  handleChange = e => {
    this.setState({ [e.target.name]: e.target.value });
  };

  handleResponseData = (data) => {
    const vCardResponse = newVCardResponse(data);
    if (vCardResponse.success) {
      this.setState({vcard_text: vCardResponse.vcard_text});
      this.setState({png_base64: vCardResponse.png_base64});
      this.setState({hide: true});
    } else {
      this.setState({error: vCardResponse.errors});
      this.setState({hide: false});
    }
  };

  onSubmit = e => {
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
    axios.post('/api/v1/vcard/create', vCardRequest)
      .then(response => {
        handleResponseData(response.data);
      }).catch(error => {
        console.log("error:", error);
        handleResponseData(error.response.data);
      });
  };

  render() {
    const formClasses = ['contactqr-form-component'];
    const { first, last, company_name, title, email, cell_phone,
      street, city, state, postal_code, facebook_url, twitter_handle,
      url, note, vcard_text, error, png_base64 } = this.state;

      // show error message
    let showError = false;
    if (error.length) {
      showError = true;
    }

    // show the vCard and QR Code
    let showData = false;
    if (png_base64.length > 0) {
      showData = true;
      formClasses.push('hide');
    }

    return (
      <div>

        <p>
          Make sharing your contact information easy. Create a QR code with your contact information that can be scanned with any phone, instead of sending someone a SMS when you meet them.
        </p>

        <form className={formClasses.join(' ')} name="contactQRForm" onSubmit={this.onSubmit} noValidate>

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
                <input type="text" className="form-control" placeholder="blackstar" name="twitter_handle" value={twitter_handle} onChange={this.handleChange} />
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

          <ErrorMessage text={error} show={showError} />

          <button className="btn btn-primary" type="submit">Create</button>
        </form>

        <CreateAnotherButton onClickAction={this.resetFormState} show={showData} />

        <VCardQRCode png_base64={png_base64} vcard_text={vcard_text} show={showData} />

      </div>
    );
  }
}