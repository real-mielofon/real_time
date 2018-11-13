import React, { Component } from 'react';
import PropTypes from 'prop-types';

import ChannelForm from './ChannelForm.jsx';
import ChannelList  from './ChannelList.jsx';


class ChannelSection extends Component{
        render() {
            return (
                <div>
                    <ChannelList {...this.props} />
                    <ChannelForm {...this.props} />
                </div>
            )
        }
}

// ChannelSection.propTypes = {
//     channels: PropTypes.array.isRequred,
//     setChannel: PropTypes.func.isRequred,
//     addChannel: PropTypes.func.isRequred
// }

export default ChannelSection;