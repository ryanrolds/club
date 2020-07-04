import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import Room from '../room'

const App = ({ props }) => {
  return (
    <div>
      <CssBaseline />
      <TopBar />
      <Room />
    </div>
  )
}

const mapStateToProps = (state) => {
  return {
    connected: state.connected,
    group: state.group
  }
}


App.propTypes = {
  connected :PropTypes.string.isRequired,
  group: PropTypes.shape({}).isRequired,
}

export default connect(mapStateToProps)(App);