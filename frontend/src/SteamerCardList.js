import React from 'react'
import { makeStyles, Container, Grid } from '@material-ui/core'
import StreamerCard from './StreamerCard'
import SignallingServer from './helpers/signalling'

export default class StreamerCardList extends React.Component {
    constructor(props) {
        super(props)
        this.streamersList = []
        this.peers = {}
        this.signals = new SignallingServer()
        this.stream = null
        this.offer = null
    }

    useStyles() {
        return makeStyles((theme) => ({
            cardGrid: {
                paddingTop: theme.spacing(1),
                paddingBottom: theme.spacing(1),
            },
        }));
    }

    getContainerClasses() {
        const classes = this.useStyles();
        return classes.cardGrid 
    }

    render() {
        return (
            <Container className={this.getContainerClasses()}>
                <Grid container spacing={1}>
                    {this.streamersList.length && this.streamersList.map((stream) => (
                        <Grid item key={stream} xs={12} sm={6}>
                            <StreamerCard stream={stream} />
                        </Grid>
                    ))}
                </Grid>
            </Container>
        )
    }

    componentDidMount() {
        this.setupMedia()
        this.updateStreamersList()
    }

    componentDidUpdate() {
        this.updateStreamersList()
    }

    async setupMedia() {
        const opts = { audio: true, video: true }
        const stream = await navigator.mediaDevices.getUserMedia(opts)
        this.setState({ stream })
        this.setState({ streamersList: this.streamersList.push(stream) })
    }

    updateStreamersList() {

    }
}