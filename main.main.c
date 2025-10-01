
/* WARNING: Unknown calling convention -- yet parameter storage is locked */

void main_main(void)

{
  cosine_sh_cli_config_CosineCliConfig *pcVar1;
  bool bVar2;
  github_com_urfave_cli_v3_Command *c;
  github_com_urfave_cli_v3_Command *pgVar3;
  github_com_urfave_cli_v3_Command *pgVar4;
  github_com_urfave_cli_v3_Command *pgVar5;
  github_com_urfave_cli_v3_Command *pgVar6;
  github_com_urfave_cli_v3_Command *pgVar7;
  github_com_urfave_cli_v3_Command *pgVar8;
  undefined8 *puVar9;
  github_com_urfave_cli_v3_Command *pgVar10;
  github_com_urfave_cli_v3_Command *pgVar11;
  github_com_urfave_cli_v3_Command *pgVar12;
  long lVar13;
  void *pvVar14;
  github_com_urfave_cli_v3_Command **ppgVar15;
  undefined1 *puVar16;
  github_com_urfave_cli_v3_Command *in_R11;
  long unaff_R14;
  undefined **in_XMM15_Qa;
  int in_XMM15_Qb;
  string sVar17;
  error eVar18;
  string format;
  context_Context ctx;
  __string osArgs;
  string key;
  __interface___ v;
  undefined1 auVar19 [16];
  undefined1 auVar20 [16];
  undefined1 auVar21 [16];
  undefined1 auVar22 [16];
  context_backgroundCtx _autotmp_56;
  github_com_getsentry_sentry_go_ClientOptions in_stack_fffffffffffff9d8;
  github_com_urfave_cli_v3_Command *local_448;
  undefined *local_440;
  undefined8 uStack_438;
  undefined *local_418;
  undefined8 local_410;
  undefined8 *local_390;
  undefined8 local_388;
  undefined8 local_380;
  undefined8 *local_340;
  undefined1 local_1e8 [32];
  undefined *local_1c8;
  undefined8 local_1c0;
  uint8 *local_100;
  int local_f8;
  uint8 *local_e0;
  int local_d8;
  internal_runtime_maps_Map *local_78;
  internal_abi_Type *local_68;
  void *pvStack_60;
  github_com_urfave_cli_v3_Command *local_58;
  github_com_urfave_cli_v3_Command *local_50;
  github_com_urfave_cli_v3_Command *local_48;
  github_com_urfave_cli_v3_Command *local_40;
  github_com_urfave_cli_v3_Command *local_38;
  github_com_urfave_cli_v3_Command *local_30;
  github_com_urfave_cli_v3_Command *local_28;
  internal_runtime_maps_Map *local_20;
  github_com_urfave_cli_v3_Command *local_18;
  undefined **local_10;
  
                    /* Unresolved local var: cosine.sh/cli/config.Config * conf@[???]
                       Unresolved local var: cosine.sh/cli/cmd.LoginCommand * login@[???]
                       Unresolved local var: cosine.sh/cli/cmd.LogoutCommand * logout@[???]
                       Unresolved local var: cosine.sh/cli/cmd.ServeCommand * serve@[???]
                       Unresolved local var: cosine.sh/cli/cmd.StartCommand * start@[???]
                       Unresolved local var: cosine.sh/cli/cmd.DiffCommand * diff@[???]
                       Unresolved local var: cosine.sh/cli/cmd.TerminalSetupCommand *
                       termSetup@[???]
                       Unresolved local var: cosine.sh/cli/cmd.DaemonCommand * daemon@[???]
                       Unresolved local var: github.com/urfave/cli/v3.Command * app@[???] */
  while (&stack0xfffffffffffffa58 <= *(undefined1 **)(unaff_R14 + 0x10)) {
    runtime_morestack_noctxt();
  }
  local_10 = in_XMM15_Qa;
  sVar17 = cosine_sh_cli_config_getConfigFile();
  local_18 = (github_com_urfave_cli_v3_Command *)sVar17.str;
  c = runtime_newobject((internal_abi_Type *)&DAT_00ee2480);
  (c->Name).len = sVar17.len;
  if (runtime_writeBarrier._0_4_ != 0) {
    c = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_18;
  }
  (c->Name).str = (uint8 *)local_18;
  pgVar3 = runtime_newobject((internal_abi_Type *)&DAT_00f3d520);
  *(bool *)&(pgVar3->ArgsUsage).str = true;
  ((struct___AutoAccept_bool__toml___auto_accept_____SystemPromptId_string__toml___system_prompt_id_____
    *)&(pgVar3->ArgsUsage).len)->AutoAccept = false;
  (pgVar3->Version).str = (uint8 *)in_XMM15_Qa;
  (pgVar3->Version).len = in_XMM15_Qb;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar3 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier2();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar3;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) =
         (github_com_urfave_cli_v3_Command *)(c->Aliases).array;
  }
  (c->Aliases).array = &pgVar3->Name;
  cosine_sh_cli_config___Config__Load((cosine_sh_cli_config_Config *)c);
  if ((((cosine_sh_cli_config_CosineCliConfig *)(c->Aliases).array)->Telemetry).Enabled == false) {
    bVar2 = false;
  }
  else {
                    /* Unresolved local var: error err@[???] */
    pgVar3 = c;
    FUN_0047e16f(local_1e8);
    local_1c8 = &DAT_010100f0;
    local_1c0 = 0x5f;
    pcVar1 = (cosine_sh_cli_config_CosineCliConfig *)(pgVar3->Aliases).array;
    local_e0 = (pcVar1->Backend).Host.str;
    local_d8 = (pcVar1->Backend).Host.len;
    local_100 = main_Version.str;
    local_f8 = main_Version.len;
    local_20 = runtime_makemap_small();
    sVar17.len = 2;
    sVar17.str = &DAT_00fbf918;
    puVar9 = runtime_mapassign_faststr((internal_abi_SwissMapType *)&DAT_00e9c300,local_20,sVar17);
    puVar9[1] = 5;
    if (runtime_writeBarrier._0_4_ != 0) {
      auVar19 = runtime_gcWriteBarrier1();
      puVar9 = auVar19._0_8_;
      *(github_com_urfave_cli_v3_Command **)in_R11 = auVar19._8_8_;
    }
    *puVar9 = &DAT_00fc1bd6;
    key.len = 4;
    key.str = &DAT_00fc0962;
    puVar9 = runtime_mapassign_faststr((internal_abi_SwissMapType *)&DAT_00e9c300,local_20,key);
    puVar9[1] = 5;
    if (runtime_writeBarrier._0_4_ != 0) {
      pgVar3 = (github_com_urfave_cli_v3_Command *)*puVar9;
      puVar9 = (undefined8 *)runtime_gcWriteBarrier1();
      *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar3;
    }
    *puVar9 = &DAT_00fc1bdb;
    local_78 = local_20;
    FUN_0047e44c(&stack0xfffffffffffff9d8,&local_1c8);
    eVar18 = github_com_getsentry_sentry_go_Init(in_stack_fffffffffffff9d8);
    if (eVar18.tab != (internal_abi_ITab *)0x0) {
      local_68 = (eVar18.tab)->Type;
      format.len = 0xf;
      format.str = &DAT_00fd5a1b;
      v.len = 1;
      v.array = (interface___ *)&local_68;
      v.cap = 1;
      pvStack_60 = eVar18.data;
      log_Fatalf(format,v);
    }
    local_10 = &PTR_main_main_deferwrap1_0123f3a8;
    bVar2 = true;
  }
  pgVar3 = runtime_newobject((internal_abi_Type *)&DAT_00ea89e0);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar3 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar3->Name).str = (uint8 *)c;
  pgVar4 = runtime_newobject((internal_abi_Type *)&DAT_00ea8a60);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar4 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar4->Name).str = (uint8 *)c;
  pgVar5 = runtime_newobject((internal_abi_Type *)&DAT_00ea8ae0);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar5 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar5->Name).str = (uint8 *)c;
  pgVar6 = runtime_newobject((internal_abi_Type *)&DAT_00ea8b60);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar6 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar6->Name).str = (uint8 *)c;
  pgVar7 = runtime_newobject((internal_abi_Type *)&DAT_00ea8960);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar7 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar7->Name).str = (uint8 *)c;
  pgVar8 = runtime_newobject((internal_abi_Type *)&DAT_00ea88e0);
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar8 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
  }
  (pgVar8->Name).str = (uint8 *)c;
  local_28 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_28->Name).len = 3;
  (local_28->Name).str = &DAT_00fbfe79;
  (local_28->UsageText).len = 0x3c;
  (local_28->UsageText).str = &DAT_0100886e;
  (local_28->Description).len = 1;
  (local_28->Description).str = &DAT_014160e0;
  *(undefined1 *)&(local_28->Version).len = 0;
  local_30 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_30->Name).len = 7;
  (local_30->Name).str = &DAT_00fc551c;
  (local_30->UsageText).len = 0x12;
  (local_30->UsageText).str = &DAT_00fdb23f;
  *(undefined1 *)&(local_30->Version).len = 0;
  local_38 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_38->Name).len = 8;
  (local_38->Name).str = &DAT_00fc8034;
  (local_38->UsageText).len = 0x33;
  (local_38->UsageText).str = &DAT_010034ca;
  *(undefined1 *)&(local_38->Version).len = 0;
  local_40 = runtime_newobject((internal_abi_Type *)&DAT_00fa2f20);
  (local_40->Name).len = 6;
  (local_40->Name).str = (uint8 *)0xfc367a;
  puVar9 = runtime_newobject((internal_abi_Type *)&DAT_00e367e0);
  auVar19._8_8_ = local_40;
  auVar19._0_8_ = puVar9;
  puVar9[1] = 1;
  *puVar9 = &DAT_0141a4e0;
  (local_40->DefaultCommand).str = (uint8 *)0x1;
  (local_40->DefaultCommand).len = 1;
  if (runtime_writeBarrier._0_4_ != 0) {
    auVar19 = runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = auVar19._0_8_;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) =
         *(github_com_urfave_cli_v3_Command **)(auVar19._8_8_ + 0x70);
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         *(github_com_urfave_cli_v3_Command **)(auVar19._8_8_ + 0x38);
  }
  lVar13 = auVar19._8_8_;
  *(long *)(lVar13 + 0x70) = auVar19._0_8_;
  *(undefined8 *)(lVar13 + 0x40) = 0x3b;
  *(undefined **)(lVar13 + 0x38) = &DAT_01007fe7;
  *(undefined1 *)(lVar13 + 99) = 0;
  FUN_0047e10e(&local_448);
  uStack_438 = 5;
  local_440 = &DAT_00fc1be0;
  local_410 = 0x20;
  local_418 = &DAT_00fefdb1;
  local_390 = runtime_newobject((internal_abi_Type *)&DAT_00e36900);
  puVar16 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  *local_390 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    local_390 = (undefined8 *)runtime_gcWriteBarrier4();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_28;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_30;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = local_38;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = local_40;
  }
  local_390[1] = local_28;
  local_390[2] = puVar16;
  local_390[3] = local_30;
  local_390[4] = puVar16;
  local_390[5] = local_38;
  local_390[6] = 
  go_itab__github_com_urfave_cli_v3_FlagBase_bool_github_com_urfave_cli_v3_BoolConfig_github_com_urfave_cli_v3_boolValue__github_com_urfave_cli_v3_Flag
  ;
  local_390[7] = local_40;
  local_388 = 4;
  local_380 = 4;
  local_340 = runtime_newobject((internal_abi_Type *)&DAT_00ebeb20);
  *local_340 = cosine_sh_cli_cmd___StartCommand__RunTeaApp_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    local_340 = (undefined8 *)runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar6;
  }
  local_340[1] = pgVar6;
  pgVar10 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar10->Name).len = 5;
  (pgVar10->Name).str = &DAT_00fc1be5;
  (pgVar10->Usage).len = 0x1c;
  (pgVar10->Usage).str = &DAT_00feab23;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebe9a0);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___LoginCommand__Login_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar3;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar10->Action;
  }
  (pgVar11->Name).len = (int)pgVar3;
  pgVar10->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  pgVar3 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar3->Name).len = 6;
  (pgVar3->Name).str = (uint8 *)0xfc3680;
  (pgVar3->Usage).len = 0x1f;
  (pgVar3->Usage).str = &DAT_00feea32;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebea20);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___LogoutCommand__Logout_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar4;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar3->Action;
  }
  (pgVar11->Name).len = (int)pgVar4;
  pgVar3->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  local_28 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_28->Name).len = 3;
  (local_28->Name).str = &DAT_00fbfe79;
  (local_28->UsageText).len = 0x3c;
  (local_28->UsageText).str = &DAT_0100886e;
  (local_28->Description).len = 1;
  (local_28->Description).str = &DAT_014160e0;
  *(undefined1 *)&(local_28->Version).len = 0;
  local_30 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_30->Name).len = 7;
  (local_30->Name).str = &DAT_00fc551c;
  (local_30->UsageText).len = 0x12;
  (local_30->UsageText).str = &DAT_00fdb23f;
  *(undefined1 *)&(local_30->Version).len = 1;
  local_40 = runtime_newobject((internal_abi_Type *)&DAT_00fa2f20);
  (local_40->Name).len = 7;
  (local_40->Name).str = &DAT_00fc5523;
  (local_40->UsageText).len = 0x17;
  (local_40->UsageText).str = &DAT_00fe3b8e;
  *(undefined1 *)&(local_40->Version).len = 0;
  pgVar4 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar4->Name).len = 5;
  (pgVar4->Name).str = &DAT_00fc1bea;
  (pgVar4->Usage).len = 0x1b;
  (pgVar4->Usage).str = &DAT_00fe974f;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e368a0);
  puVar16 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->Name).str =
       go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier5();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_28;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_30;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = local_40;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x20) =
         (github_com_urfave_cli_v3_Command *)(pgVar4->Flags).array;
  }
  (pgVar11->Name).len = (int)local_28;
  (pgVar11->Aliases).array = (string *)puVar16;
  (pgVar11->Aliases).len = (int)local_30;
  (pgVar11->Aliases).cap =
       (int)
       go_itab__github_com_urfave_cli_v3_FlagBase_bool_github_com_urfave_cli_v3_BoolConfig_github_com_urfave_cli_v3_boolValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->Usage).str = (uint8 *)local_40;
  (pgVar4->Flags).len = 3;
  (pgVar4->Flags).cap = 3;
  (pgVar4->Flags).array = (github_com_urfave_cli_v3_Flag *)pgVar11;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebeaa0);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___ServeCommand__Serve_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar5;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar4->Action;
  }
  (pgVar11->Name).len = (int)pgVar5;
  pgVar4->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  local_28 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_28->Name).len = 3;
  (local_28->Name).str = &DAT_00fbfe79;
  (local_28->UsageText).len = 0x3c;
  (local_28->UsageText).str = &DAT_0100886e;
  (local_28->Description).len = 1;
  (local_28->Description).str = &DAT_014160e0;
  *(undefined1 *)&(local_28->Version).len = 0;
  local_30 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_30->Name).len = 7;
  (local_30->Name).str = &DAT_00fc551c;
  (local_30->UsageText).len = 0x12;
  (local_30->UsageText).str = &DAT_00fdb23f;
  *(undefined1 *)&(local_30->Version).len = 0;
  local_38 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_38->Name).len = 8;
  (local_38->Name).str = &DAT_00fc8034;
  (local_38->UsageText).len = 0x33;
  (local_38->UsageText).str = &DAT_010034ca;
  *(undefined1 *)&(local_38->Version).len = 0;
  local_40 = runtime_newobject((internal_abi_Type *)&DAT_00fa2f20);
  (local_40->Name).len = 6;
  (local_40->Name).str = (uint8 *)0xfc367a;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e367e0);
  (pgVar11->Name).len = 1;
  (pgVar11->Name).str = &DAT_0141a4e0;
  (local_40->DefaultCommand).str = (uint8 *)0x1;
  (local_40->DefaultCommand).len = 1;
  pgVar5 = local_40;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) =
         (github_com_urfave_cli_v3_Command *)(pgVar5->Description).len;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)(pgVar5->UsageText).str;
  }
  (pgVar5->Description).len = (int)pgVar11;
  (pgVar5->UsageText).len = 0x3b;
  (pgVar5->UsageText).str = &DAT_01007fe7;
  *(undefined1 *)((long)&(pgVar5->Version).len + 3) = 0;
  pgVar5 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar5->Name).len = 5;
  (pgVar5->Name).str = &DAT_00fc1be0;
  (pgVar5->Usage).len = 0x20;
  (pgVar5->Usage).str = &DAT_00fefdb1;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e36900);
  puVar16 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->Name).str =
       go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier6();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_28;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_30;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = local_38;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = local_40;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x20) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x28) =
         (github_com_urfave_cli_v3_Command *)(pgVar5->Flags).array;
  }
  (pgVar11->Name).len = (int)local_28;
  (pgVar11->Aliases).array = (string *)puVar16;
  (pgVar11->Aliases).len = (int)local_30;
  (pgVar11->Aliases).cap = (int)puVar16;
  (pgVar11->Usage).str = (uint8 *)local_38;
  (pgVar11->Usage).len =
       (int)
       go_itab__github_com_urfave_cli_v3_FlagBase_bool_github_com_urfave_cli_v3_BoolConfig_github_com_urfave_cli_v3_boolValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->UsageText).str = (uint8 *)local_40;
  (pgVar5->Flags).len = 4;
  (pgVar5->Flags).cap = 4;
  (pgVar5->Flags).array = (github_com_urfave_cli_v3_Flag *)pgVar11;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebeb20);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___StartCommand__RunTeaApp_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar6;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar5->Action;
  }
  (pgVar11->Name).len = (int)pgVar6;
  pgVar5->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  local_28 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_28->Name).len = 5;
  (local_28->Name).str = &DAT_00fc1bef;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e367e0);
  (pgVar11->Name).len = 1;
  (pgVar11->Name).str = &go_string__;
  (local_28->Category).str = (uint8 *)0x1;
  (local_28->Category).len = 1;
  pgVar6 = local_28;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) =
         (github_com_urfave_cli_v3_Command *)(pgVar6->DefaultCommand).len;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)(pgVar6->UsageText).str;
  }
  (pgVar6->DefaultCommand).len = (int)pgVar11;
  (pgVar6->UsageText).len = 0x33;
  (pgVar6->UsageText).str = &DAT_010034fd;
  *(undefined1 *)&(pgVar6->Version).len = 0;
  local_30 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_30->Name).len = 4;
  (local_30->Name).str = &DAT_00fc0966;
  (local_30->UsageText).len = 0x2b;
  (local_30->UsageText).str = &DAT_00ffc63a;
  *(undefined1 *)&(local_30->Version).len = 0;
  local_38 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_38->Name).len = 3;
  (local_38->Name).str = &DAT_00fbfe79;
  (local_38->UsageText).len = 0x4b;
  (local_38->UsageText).str = &DAT_0100d58c;
  *(undefined1 *)&(local_38->Version).len = 0;
  pgVar6 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar6->Name).len = 4;
  (pgVar6->Name).str = &DAT_00fc096a;
  (pgVar6->Usage).len = 0x32;
  (pgVar6->Usage).str = &DAT_0100287c;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e368a0);
  puVar16 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->Name).str =
       go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier5();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_28;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_30;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = local_38;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x20) =
         (github_com_urfave_cli_v3_Command *)(pgVar6->Flags).array;
  }
  (pgVar11->Name).len = (int)local_28;
  (pgVar11->Aliases).array = (string *)puVar16;
  (pgVar11->Aliases).len = (int)local_30;
  (pgVar11->Aliases).cap = (int)puVar16;
  (pgVar11->Usage).str = (uint8 *)local_38;
  (pgVar6->Flags).len = 3;
  (pgVar6->Flags).cap = 3;
  (pgVar6->Flags).array = (github_com_urfave_cli_v3_Flag *)pgVar11;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebe920);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___DiffCommand__Run_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar7;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar6->Action;
  }
  (pgVar11->Name).len = (int)pgVar7;
  pgVar6->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  local_40 = runtime_newobject((internal_abi_Type *)&DAT_00fa2f20);
  (local_40->Name).len = 7;
  (local_40->Name).str = &DAT_00fc552a;
  (local_40->UsageText).len = 0x33;
  (local_40->UsageText).str = &DAT_01003530;
  *(undefined1 *)((long)&(local_40->Version).len + 3) = 0;
  local_48 = runtime_newobject((internal_abi_Type *)&DAT_00fa2f20);
  (local_48->Name).len = 5;
  (local_48->Name).str = &DAT_00fc1bf4;
  (local_48->UsageText).len = 0x29;
  (local_48->UsageText).str = &DAT_00ffa71a;
  *(undefined1 *)((long)&(local_48->Version).len + 3) = 0;
  pgVar7 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar7->Name).len = 0xe;
  (pgVar7->Name).str = &DAT_00fd3edf;
  (pgVar7->Usage).len = 0x3b;
  (pgVar7->Usage).str = &DAT_01008022;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00e36840);
  puVar16 = 
  go_itab__github_com_urfave_cli_v3_FlagBase_bool_github_com_urfave_cli_v3_BoolConfig_github_com_urfave_cli_v3_boolValue__github_com_urfave_cli_v3_Flag
  ;
  (pgVar11->Name).str =
       go_itab__github_com_urfave_cli_v3_FlagBase_bool_github_com_urfave_cli_v3_BoolConfig_github_com_urfave_cli_v3_boolValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier4();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_40;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_48;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) =
         (github_com_urfave_cli_v3_Command *)(pgVar7->Flags).array;
  }
  (pgVar11->Name).len = (int)local_40;
  (pgVar11->Aliases).array = (string *)puVar16;
  (pgVar11->Aliases).len = (int)local_48;
  (pgVar7->Flags).len = 2;
  (pgVar7->Flags).cap = 2;
  (pgVar7->Flags).array = (github_com_urfave_cli_v3_Flag *)pgVar11;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00ebeba0);
  (pgVar11->Name).str = (uint8 *)cosine_sh_cli_cmd___TerminalSetupCommand__Run_fm;
  (pgVar11->Name).len = (int)&runtime_zerobase;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar11 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier2();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) =
         (github_com_urfave_cli_v3_Command *)pgVar7->Action;
  }
  pgVar7->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar11;
  local_28 = runtime_newobject((internal_abi_Type *)&DAT_00fa2c80);
  (local_28->Name).len = 3;
  (local_28->Name).str = &DAT_00fbfe79;
  (local_28->UsageText).len = 0x3c;
  (local_28->UsageText).str = &DAT_0100886e;
  *(undefined1 *)&(local_28->Version).len = 1;
  local_50 = runtime_newobject((internal_abi_Type *)&DAT_00fa31c0);
  (local_50->Name).len = 5;
  (local_50->Name).str = &DAT_00fc1bf9;
  (local_50->UsageText).len = 0x1c;
  (local_50->UsageText).str = &DAT_00feab3f;
  *(undefined1 *)&(local_50->Version).len = 0;
  local_58 = runtime_newobject((internal_abi_Type *)&DAT_00fa31c0);
  (local_58->Name).len = 4;
  (local_58->Name).str = &DAT_00fc096e;
  (local_58->UsageText).len = 0x21;
  (local_58->UsageText).str = &DAT_00ff14b1;
  *(undefined4 *)((long)&(local_58->Version).len + 4) = 0x1b59;
  *(undefined1 *)&(local_58->Version).len = 0;
  pgVar11 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  (pgVar11->Name).len = 6;
  (pgVar11->Name).str = (uint8 *)0xfc3686;
  (pgVar11->Usage).len = 0x1d;
  (pgVar11->Usage).str = &DAT_00fec093;
  pgVar11->Hidden = true;
  pgVar12 = runtime_newobject((internal_abi_Type *)&DAT_00e368a0);
  (pgVar12->Name).str =
       go_itab__github_com_urfave_cli_v3_FlagBase_string_github_com_urfave_cli_v3_StringConfig_github_com_urfave_cli_v3_stringValue__github_com_urfave_cli_v3_Flag
  ;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar12 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier5();
    *(github_com_urfave_cli_v3_Command **)in_R11 = local_28;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = local_50;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = local_58;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = pgVar12;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x20) =
         (github_com_urfave_cli_v3_Command *)(pgVar11->Flags).array;
  }
  (pgVar12->Name).len = (int)local_28;
  (pgVar12->Aliases).array =
       (string *)
       go_itab__github_com_urfave_cli_v3_FlagBase_int32_github_com_urfave_cli_v3_IntegerConfig_github_com_urfave_cli_v3_intValue_int32___github_com_urfave_cli_v3_Flag
  ;
  (pgVar12->Aliases).len = (int)local_50;
  (pgVar12->Aliases).cap =
       (int)
       go_itab__github_com_urfave_cli_v3_FlagBase_int32_github_com_urfave_cli_v3_IntegerConfig_github_com_urfave_cli_v3_intValue_int32___github_com_urfave_cli_v3_Flag
  ;
  (pgVar12->Usage).str = (uint8 *)local_58;
  (pgVar11->Flags).len = 3;
  (pgVar11->Flags).cap = 3;
  (pgVar11->Flags).array = (github_com_urfave_cli_v3_Flag *)pgVar12;
  pgVar12 = runtime_newobject((internal_abi_Type *)&DAT_00ebe8a0);
  (pgVar12->Name).str = (uint8 *)cosine_sh_cli_cmd___DaemonCommand__ServeGrpcServer_fm;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar12 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar8;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar12;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)pgVar11->Action;
  }
  (pgVar12->Name).len = (int)pgVar8;
  pgVar11->Action = (github_com_urfave_cli_v3_ActionFunc **)pgVar12;
  puVar9 = runtime_newobject((internal_abi_Type *)&DAT_00fb8ae0);
  puVar9[1] = 3;
  *puVar9 = &DAT_00fbfe7c;
  puVar9[6] = 0x31;
  puVar9[5] = &DAT_01001a83;
  auVar20._8_8_ = main_Version.str;
  auVar20._0_8_ = puVar9;
  puVar9[0xc] = main_Version.len;
  if (runtime_writeBarrier._0_4_ != 0) {
    auVar20 = runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = auVar20._8_8_;
  }
  lVar13 = auVar20._0_8_;
  auVar21._8_8_ = local_440;
  auVar21._0_8_ = lVar13;
  *(long *)(lVar13 + 0x58) = auVar20._8_8_;
  *(undefined8 *)(lVar13 + 0x80) = uStack_438;
  if (runtime_writeBarrier._0_4_ != 0) {
    auVar21 = runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = auVar21._8_8_;
  }
  local_448 = auVar21._0_8_;
  (local_448->DefaultCommand).str = auVar21._8_8_;
  pvVar14 = runtime_newobject((internal_abi_Type *)&DAT_00e36960);
  auVar22._8_8_ = pgVar3;
  auVar22._0_8_ = pvVar14;
  pgVar8 = local_448;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar8 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier8();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar10;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar3;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) = pgVar4;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x18) = pgVar5;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x20) = pgVar6;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x28) = pgVar7;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x30) = pgVar11;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x38) = pgVar8;
    pgVar3 = (github_com_urfave_cli_v3_Command *)(local_448->Commands).array;
    pgVar8 = local_448;
    in_R11 = pgVar3;
    auVar22 = runtime_gcWriteBarrier1();
    *(github_com_urfave_cli_v3_Command **)in_R11 = pgVar3;
  }
  ppgVar15 = auVar22._0_8_;
  *ppgVar15 = pgVar10;
  ppgVar15[1] = auVar22._8_8_;
  ppgVar15[2] = pgVar4;
  ppgVar15[3] = pgVar5;
  ppgVar15[4] = pgVar6;
  ppgVar15[5] = pgVar7;
  ppgVar15[6] = pgVar11;
  (pgVar8->Commands).len = 7;
  (pgVar8->Commands).cap = 7;
  (pgVar8->Commands).array = ppgVar15;
  pgVar3 = runtime_newobject((internal_abi_Type *)&DAT_00ebec20);
  (pgVar3->Name).str = (uint8 *)main_main_func1;
  if (runtime_writeBarrier._0_4_ != 0) {
    pgVar3 = (github_com_urfave_cli_v3_Command *)runtime_gcWriteBarrier3();
    *(github_com_urfave_cli_v3_Command **)in_R11 = c;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 8) = pgVar3;
    *(github_com_urfave_cli_v3_Command **)((long)in_R11 + 0x10) =
         (github_com_urfave_cli_v3_Command *)local_448->Before;
  }
  (pgVar3->Name).len = (int)c;
  local_448->Before = (github_com_urfave_cli_v3_BeforeFunc **)pgVar3;
  ctx.data = &runtime_zerobase;
  ctx.tab = (internal_abi_ITab *)go_itab_context_backgroundCtx_context_Context;
  osArgs.len = os_Args.len;
  osArgs.array = os_Args.array;
  osArgs.cap = os_Args.cap;
  github_com_urfave_cli_v3___Command__run(local_448,ctx,osArgs);
  if (bVar2) {
    main_main_deferwrap1();
  }
  return;
}

