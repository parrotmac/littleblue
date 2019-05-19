import React, { Component } from 'react';

import { Card, Elevation, H5 } from '@blueprintjs/core';

import './SettingsPage.scss';

class SettingsPage extends Component {
    public render(): JSX.Element {
        return (
            <div className={'SettingsPage'}>
                <Card elevation={Elevation.TWO}>
                    <H5>Build Settings</H5>
                    <p>Nothing here, yet.</p>
                </Card>
            </div>
        );
    }
}

export default SettingsPage;
