import React from 'react';

import {
    Button,
    Icon,
    InputGroup,
    Intent,
    Spinner,
    Tooltip,
} from '@blueprintjs/core';

import './LoginForm.css';

export interface ILoginFormState {
    disabled: boolean;
    showSpinner: boolean;
    showPassword: boolean;
    formUsername: string;
    formPassword: string;
}

class LoginForm extends React.Component<{}, ILoginFormState> {
    public state: ILoginFormState = {
        disabled: false,
        showSpinner: false,
        showPassword: false,
        formUsername: '',
        formPassword: ''
    };

    private handleLockClick = () => {
        const newState = {
            showPassword: !this.state.showPassword,
        };
        this.setState(newState as any);
    };

    private createOnFormUpdate = (formKey: string) => {
        return (event: React.FormEvent<HTMLInputElement>) => {
            let newStateMod = {} as any;
            newStateMod[formKey] = event.currentTarget.value;
            this.setState(newStateMod);
        }
    };

    private handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const { formUsername, formPassword } = this.state;
        console.log("login:", formUsername, formPassword);
        console.warn("submit action suppressed");
        return false;
    };

    public render(): JSX.Element {

        const {
            disabled,
            showSpinner,
            showPassword,
            formUsername,
            formPassword
        } = this.state;

        const maybeSpinner = showSpinner ? <Spinner size={Icon.SIZE_STANDARD}/> : undefined;

        const lockButton = (
            <Tooltip content={`${showPassword ? 'Hide' : 'Show'} Password`} disabled={disabled}>
                <Button
                    disabled={disabled}
                    icon={showPassword ? 'unlock' : 'lock'}
                    intent={Intent.WARNING}
                    minimal={true}
                    onClick={this.handleLockClick}
                    tabIndex={-1}
                />
            </Tooltip>
        );

        return (
            <form onSubmit={this.handleSubmit} className={"LoginForm"}>
                <InputGroup
                    large={true}
                    leftIcon="user"
                    placeholder="you@example.com"
                    rightElement={maybeSpinner}
                    value={formUsername}
                    name={"email"}
                    onChange={this.createOnFormUpdate("formUsername")}
                />
                <InputGroup
                    large={true}
                    leftIcon={'key'}
                    placeholder="**********"
                    rightElement={lockButton}
                    type={showPassword ? 'text' : 'password'}
                    value={formPassword}
                    name={"password"}
                    onChange={this.createOnFormUpdate("formPassword")}
                />
                <Button
                    large={true}
                    disabled={formUsername === "" || formPassword === ""}
                    rightIcon="arrow-right"
                    intent="success"
                    text="Login"
                    type={"submit"}
                />
            </form>
        );
    }
}

export default LoginForm;
