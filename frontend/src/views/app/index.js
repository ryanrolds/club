import React, { useState } from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import GroupGridList from '../../molecules/group/groupGridList'
import { Paper, makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  root: {
    paddingTop: '16px',
  }
})

export default function App() {
  const classes = useStyles()
  const [singerData, setSingerData] = useState({ id: '123' })
  const [localData, setLocalData] = useState({ id: '9999' })
  const [peersData, setPeersData] = useState([
    { id: '1' },
    { id: '2' },
    { id: '3' },
    { id: '4' },
    { id: '5' },
  ])

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <Paper className={classes.root}>
          <GroupGridList singer={singerData} local={localData} peers={peersData} />
        </Paper>
      </main>
    </>
  )
}
