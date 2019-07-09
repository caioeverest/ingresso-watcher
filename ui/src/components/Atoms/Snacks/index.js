import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components'
import { amber, green, red, blue } from '@material-ui/core/colors';
import { Snackbar,
    SnackbarContent,
} from '@material-ui/core';
import { ErrorIcon,
    InfoIcon,
    CheckCircleIcon,
    WarningIcon,
} from '@material-ui/icons';

const options = {
    success: {
        color: green[600],
        icon: CheckCircleIcon,
    },
    error: {
        color: red[700],
        icon: ErrorIcon,
    },
    info: {
        color: blue[600],
        icon: InfoIcon,
    },
    warning: {
        color: amber[700],
        icon: WarningIcon,
    },
};

const LINK = props => <a { ...props } />

const Content = styled(LINK)`
    background-color: ${ props => options['info'].color }
`

const Snacks = props => {
  const { open, message, time, status } = props;

  return (
      <Snackbar
          anchorOrigin={{
              vertical: 'bottom',
              horizontal: 'left',
          }}
          open={open}
          autoHideDuration={time*1000}
      >
          <SnackbarContent
              message={ message }
              classes={ `background-color: ${options['info'].color}` }
          />
      </Snackbar>
  );
}

Snacks.propTypes = {
    message: PropTypes.node,
    open: PropTypes.bool,
    timeout: PropTypes.number,
    status: PropTypes.oneOf(['success', 'warning', 'error', 'info']).isRequired,
};

export default Snacks
