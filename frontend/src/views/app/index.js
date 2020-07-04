import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import Room from '../room'
import Group from '../group'

const App = ({ connected, group }) => {
  return (
    <div>
      <CssBaseline />
      {connected !== 'connected' && <span>Connecting...</span>}
      {connected === 'connected' && <TopBar />}
      {connected === 'connected' && group.id === undefined && <Room />}
      {connected === 'connected' && group.id !== undefined && <Group />}
    </div>
  )
}

App.propTypes = {
  connected: PropTypes.string.isRequired,
  group: PropTypes.shape({
    id: PropTypes.string,
  }).isRequired,
}

const mapStateToProps = (state) => {
  return {
    connected: state.connected,
    group: state.group,
  }
}

export default connect(mapStateToProps)(App)
