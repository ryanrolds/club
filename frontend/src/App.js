import React from 'react'
import GithubIcon from '@material-ui/icons/GitHub'
import CssBaseline from '@material-ui/core/CssBaseline'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import Link from '@material-ui/core/Link'
import StreamerCardList from './StreamerCardList'

const useStyles = makeStyles((theme) => ({
  icon: {
    marginRight: theme.spacing(2),
  },
  footer: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(6),
  },
}))

function Footer() {
  const classes = useStyles()

  return (
    <footer className={classes.footer}>
      <Typography variant='h6' align='center' gutterBottom>
        Help contribute
      </Typography>
      <Typography
        variant='subtitle1'
        align='center'
        color='textSecondary'
        component='p'
      >
        <Link color='inherit' href='https://github.com/ryanrolds/club'>
          <GithubIcon className={classes.icon} />
          Contribute on GitHub
        </Link>
      </Typography>
    </footer>
  )
}

export default function Club() {
  return (
    <>
      <CssBaseline />
      <main>
        <StreamerCardList />
      </main>
      <Footer />
    </>
  )
}
