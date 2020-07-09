import React, { useState } from 'react'
import PropTypes from 'prop-types'
import Lodash from 'lodash'
import { makeStyles } from '@material-ui/core/styles'
import MemberGridRow from '../member_grid_row'

const useStyles = makeStyles({
  grid: {
    display: 'grid',
    width: '100vw',
    height: '100vh',
    'grid-gap': '1rem',
    'justify-content': 'center',
    'align-content': 'center',
  },
  gridColumns: (props) => ({
    'grid-template-columns': `repeat(${props.cols}, ${120 / props.cols}vmin)`,
    'grid-template-rows': `repeat(${props.rowss}, ${120 / props.rows}vmin)`,
  }),
})

const MemberGrid = ({ localID, members }) => {
  const [localStream, setLocalStream] = useState(null)
  const numMembers = members.length
  const sqrtMembers = Math.sqrt(numMembers)
  const numColumns = Math.ceil(sqrtMembers)
  const rows = Lodash.chunk(members, numColumns)
  const classes = useStyles({ cols: numColumns, rows: rows.length })

  return (
    <div className={`${classes.grid} ${classes.gridColumns}`}>
      {rows.map((row, index) => (
        <MemberGridRow
          key={`member_row_${index}`} // eslint-disable-line react/no-array-index-key
          cols={numColumns}
          members={row}
          localID={localID}
          localStream={localStream}
          setLocalStream={setLocalStream}
        />
      ))}
    </div>
  )
}

MemberGrid.propTypes = {
  localID: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }).isRequired
  ).isRequired,
}

export default MemberGrid
