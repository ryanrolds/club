import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import { Paper, makeStyles } from '@material-ui/core'
import TopBar from '../../components/appBar/topBar'
import PersonGridList from '../../molecules/person/personGridList'
import SignalingServer from '../../helpers/signaling'
import usePeers from '../../helpers/usePeers'
import useLocalMedia from '../../helpers/useLocalMedia'

const useStyles = makeStyles({
  root: {
    paddingTop: '16px',
  },
})

const App = () => {
  const classes = useStyles()
  const local = useLocalMedia()
  const signals = new SignalingServer()
  const peers = usePeers(local, signals)

  signals.addEventListener('connected', async () => {
    signals.sendJoin()
  })

  signals.addEventListener('disconnected', async () => {
    console.log('disconnected')
  })

  const isHTTPS = window.location.protocol !== 'https:'
  signals.connect(`${isHTTPS ? 'ws' : 'wss'}://localhost:3001/room`)

  window.addEventListener('unload', () => {
    signals.sendLeave()
  })

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <Paper className={classes.root}>
          {local ? <PersonGridList local={local} peers={peers.peers} /> : null}
        </Paper>
      </main>
    </>
  )
}

export default App
