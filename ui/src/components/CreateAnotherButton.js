import React from 'react'
import './CreateAnotherButton.css';

export default class CreateAnotherButton extends React.Component {

  render() {
    const { onClickAction, show } = this.props;
    const componentClasses = ['create-another-button-component'];

    if (show) { componentClasses.push('show'); }

    return (
      <div className={componentClasses.join(' ')}>
        <button className="btn btn-primary" onClick={() => onClickAction()}>
        Create another one
        </button>
      </div>
    );
  };
}

