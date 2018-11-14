import React, { Component } from 'react'
import PropTypes from 'prop-types';

import User from './User.jsx';

class UserList extends Component{
    render() {
      return (
        <ul className="list-group">{
            this.props.users.map(user=>{
                return <User 
                    user={user.userName}
                    key={user.id}
                    {...this.props}
                    />
            })
        }</ul>
      )
    }
}

UserList.propTypes = {
     users: PropTypes.array.isRequired,
     setUserName: PropTypes.func.isRequired,
}

export default UserList;