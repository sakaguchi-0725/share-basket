import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo'
import { Module } from '@nestjs/common'
import { GraphQLModule } from '@nestjs/graphql'
import { join } from 'path'
import { PingModule } from './ping/ping.module'
import { AuthModule } from './auth/auth.module'
import { PersonalShoppingModule } from './personal-shopping/personal-shopping.module'

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: join(process.cwd(), 'src/schema.gql'),
      sortSchema: true,
      playground: process.env.NODE_ENV === 'dev',
    }),
    PingModule,
    AuthModule,
    PersonalShoppingModule,
  ],
})
export class AppModule {}
