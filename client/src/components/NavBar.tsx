import React, { Component } from 'react';

import {
    Alignment,
    Classes,
    Navbar,
    NavbarGroup,
    NavbarHeading,
} from '@blueprintjs/core';
import { Link } from 'react-router-dom';

import './NavBar.scss';

class NavBar extends Component {
    public render(): JSX.Element {

        const navLinkClassesWithIcon = (iconName: string) =>
            `${Classes.MINIMAL} ${Classes.BUTTON} ${Classes.iconClass(iconName)}`;

        return (
            <Navbar className={'NavBar'}>
                <NavbarGroup align={Alignment.LEFT}>
                    <NavbarHeading>Little Blue</NavbarHeading>
                </NavbarGroup>
                <NavbarGroup align={Alignment.RIGHT}>
                    <Link
                        className={navLinkClassesWithIcon('heat-grid')}
                        to={'/overview'}>
                        Overview
                    </Link>
                    <Link
                        className={navLinkClassesWithIcon('list')}
                        to={'/repos'}>
                        Repositories
                    </Link>
                    <Link
                        className={navLinkClassesWithIcon('cog')}
                        to={'/settings'}>
                        Settings
                    </Link>
                    <Link
                        className={navLinkClassesWithIcon('user')}
                        to={'/login'}>
                        Account
                    </Link>
                </NavbarGroup>
            </Navbar>
        );
    }
}

export default NavBar;
