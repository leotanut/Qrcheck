import React, { Component } from "react";
import axios from "axios";

class CheckOut extends Component {
  state = {
    name: "",
    status: 2
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
      status: 2
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
          <button type="submit">Check-out</button>
        </form>
      </div>
    );
  }
}

export default CheckOut;