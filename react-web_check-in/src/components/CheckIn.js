import React, { Component } from "react";
import axios from "axios";

class CheckIn extends Component {
  state = {
    name: "",
    status: 1
  };

  onTitleChange = e => {
    this.setState({
      name: e.target.value
    });
  };

  handleSubmit = e => {
    e.preventDefault();
    const data = {
      name: this.state.name,
      status: 1
    };
    axios
      .post("/users", data)
      .then(res => console.log(res))
      .catch(err => console.log(err));
  };

  render() {
    return (
      <div className="post">
        <form className="post" onSubmit={this.handleSubmit}>
          <input
            placeholder="Name" value={this.state.name}
            onChange={this.onTitleChange} required
          />
          <button type="submit">Check-in</button>
        </form>
      </div>
    );
  }
}

export default CheckIn;