import { Inject, Injectable, OnModuleInit } from '@nestjs/common'
import { LoginInput, LoginOutput } from './dto/login.dto'
import { ClientGrpc } from '@nestjs/microservices'
import { AuthServiceClient } from 'src/gen/auth/auth'
import { firstValueFrom } from 'rxjs'
import { SignUpInput, SignUpOutput } from './dto/signup.dto'
import {
  SignUpConfirmInput,
  SignUpConfirmOutput,
} from './dto/signup-confirm.dto'

@Injectable()
export class AuthService implements OnModuleInit {
  private service: AuthServiceClient

  constructor(@Inject('AUTH_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.service = this.client.getService<AuthServiceClient>('AuthService')
  }

  async login(input: LoginInput): Promise<LoginOutput> {
    const observable = this.service.login({
      email: input.email,
      password: input.password,
    })

    const response = await firstValueFrom(observable)

    return {
      accessToken: response.accessToken,
    }
  }

  async signup(input: SignUpInput): Promise<SignUpOutput> {
    await firstValueFrom(
      this.service.signUp({
        name: input.name,
        email: input.email,
        password: input.password,
      }),
    )

    return {
      message: 'ok',
    }
  }

  async signupConfirm(input: SignUpConfirmInput): Promise<SignUpConfirmOutput> {
    await firstValueFrom(
      this.service.signUpConfirm({
        email: input.email,
        confirmationCode: input.confirmationCode,
      }),
    )

    return {
      message: 'ok',
    }
  }
}
