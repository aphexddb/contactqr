import React from 'react'
import './ErrorMessage.css';

export default class ErrorMessage extends React.Component {

  render() {
    const componentClasses = ['error-component','alert', 'alert-warning'];
    const { text, show } = this.props;

    if (show) { componentClasses.push('show'); }

    return (
      <div className={componentClasses.join(' ')} role="alert">
        {text}
      </div>
    );
  }
}