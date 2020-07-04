import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import Room from '../room'

const App = ({ connected, group }) => {
  return (
    <div>
      <CssBaseline />
      <TopBar />
      {connected !== 'connected' && <span>Connecting...</span>}
      {connected === 'connected' && group.id === undefined && <Room />}
    </div>
  )
}

const mapStateToProps = (state) => {
  return {
    connected: state.connected,
    group: state.group,
  }
}

App.propTypes = {
  connected: PropTypes.string.isRequired,
  group: PropTypes.shape({
    id: PropTypes.string,
  }).isRequired,
}

export default connect(mapStateToProps)(App)
