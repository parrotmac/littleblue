import { Card } from '@blueprintjs/core';
import React, { Component } from 'react';

class NotFound extends Component {
    public render(): JSX.Element {
        return (
            <Card className={'NotFound'}>
                <code style={{ fontSize: '32px' }}>HTTP/1.1 404 Not Found</code>
                <h4>The requested resource doesn't exist.</h4>
                <p>:(</p>
            </Card>
        );
    }
}

export default NotFound;
