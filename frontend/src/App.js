import React from 'react';
import GithubIcon from '@material-ui/icons/GitHub'
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';
import StreamerCardList from './SteamerCardList'

const useStyles = makeStyles((theme) => ({
  icon: {
    marginRight: theme.spacing(2),
  },
  footer: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(6),
  },
}));

export default function Club() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <main>
        <StreamerCardList />
      </main>
      {/* Footer */}
      <footer className={classes.footer}>
        <Typography variant="h6" align="center" gutterBottom>
          Help contribute
        </Typography>
        <Typography variant="subtitle1" align="center" color="textSecondary" component="p">
          <Link color="inherit" href="https://github.com/ryanrolds/club">
            <GithubIcon className={classes.icon} />
        Contribute on GitHub
        </Link>
        </Typography>
      </footer>
      {/* End footer */}
    </React.Fragment>
  );
}