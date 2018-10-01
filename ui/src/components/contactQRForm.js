import React from 'react'
import axios from 'axios';

// parses response data into a VCardResponse
function newVCardResponse(data) {
  if (data === undefined) {
    return null;
  }
  return {
    success: data.success,
    errors: data.errors,
    vcard_text: data.vcard_text
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
    };
  }

  handleChange = e => {
    this.setState({ [e.target.name]: e.target.value });
  };

  onSubmit = (e) => {
    e.preventDefault();
    this.setState({vcard_text: ""});
    this.setState({error: ""});

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

    // TODO: remote localhost
    console.log("TODO: remote localhost");
    axios.post('http://localhost:8080/api/v1/vcard/create', vCardRequest)
      .then(response => {
        const vCardResponse = newVCardResponse(response.data);
        if (vCardResponse.success) {
          this.setState({vcard_text: vCardResponse.vcard_text});
        } else {
          this.setState({error: vCardResponse.errors});
        }
      }).catch(error => {
        const vCardResponse = newVCardResponse(error.response.data);
        if (vCardResponse.success) {
          this.setState({vcard_text: vCardResponse.vcard_text});
        } else {
          this.setState({error: vCardResponse.errors});
        }
      });
  };

  render() {
    const { first, last, note, vcard_text, error } = this.state;

    // show error message
    let errorMsg = "";
    if (error.length) {
      errorMsg =
      <span>
        <br/>{error}
      </span>;
    }

    // show the raw vCard text
    let vCardRawText = "";
    if (vcard_text.length) {
      vCardRawText =
      <div>
        <p>vCard data in <a href="https://tools.ietf.org/html/rfc6350">RFC 6350</a> format:</p>
        <pre>{vcard_text}</pre>
      </div>;
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
              Note:<br />
              <textarea name="note" value={note} onChange={this.handleChange} />
            </label>
          </p>

          <p>
            <button type="submit">Create QR Code</button>
            {errorMsg}
          </p>
        </form>

        {vCardRawText}
      </div>
    );
  }
}