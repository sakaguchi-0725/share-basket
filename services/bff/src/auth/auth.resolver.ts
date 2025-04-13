import { Args, Mutation, Resolver } from '@nestjs/graphql'
import { AuthService } from './auth.service'
import { LoginInput, LoginOutput } from './dto/login.dto'
import { SignUpInput, SignUpOutput } from './dto/signup.dto'
import {
  SignUpConfirmInput,
  SignUpConfirmOutput,
} from './dto/signup-confirm.dto'

@Resolver()
export class AuthResolver {
  constructor(private readonly authService: AuthService) {}

  @Mutation(() => LoginOutput)
  login(@Args('input') input: LoginInput): Promise<LoginOutput> {
    return this.authService.login(input)
  }

  @Mutation(() => SignUpOutput)
  signup(@Args('input') input: SignUpInput): Promise<SignUpOutput> {
    return this.authService.signup(input)
  }

  @Mutation(() => SignUpConfirmOutput)
  signUpConfirm(
    @Args('input') input: SignUpConfirmInput,
  ): Promise<SignUpConfirmOutput> {
    return this.authService.signupConfirm(input)
  }
}
