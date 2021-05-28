/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

use super::big::NLEN;
use crate::arch::Chunk;
use crate::types::{CurvePairingType, CurveType, ModType, SexticTwist, SignOfX};

// Base Bits= 60

// Base Bits= 60
// hifive Modulus
pub const MODULUS: [Chunk; NLEN] = [
    0xFFFFFFFFFFFFFFD,
    0xFFFFFFFFFFFFFFF,
    0xFFFFFFFFFFFFFFF,
    0xFFFFFFFFFFFFFFF,
    0xFFFFFFFFFFFFFFF,
    0xFFFFFFFFF,
];
pub const R2MODP: [Chunk; NLEN] = [0x9000000000000, 0x0, 0x0, 0x0, 0x0, 0x0];
pub const MCONST: Chunk = 0x3;

// hifive Curve
pub const CURVE_COF_I: isize = 8;
pub const CURVE_A: isize = 1;
pub const CURVE_B_I: isize = 11111;
pub const CURVE_COF: [Chunk; NLEN] = [0x8, 0x0, 0x0, 0x0, 0x0, 0x0];
pub const CURVE_B: [Chunk; NLEN] = [0x2B67, 0x0, 0x0, 0x0, 0x0, 0x0];
pub const CURVE_ORDER: [Chunk; NLEN] = [
    0xB2F95973E9FA805,
    0xC0BD6B87F93BAA7,
    0x71415FA9850,
    0x0,
    0x0,
    0x200000000,
];
pub const CURVE_GX: [Chunk; NLEN] = [0xC, 0x0, 0x0, 0x0, 0x0, 0x0];
pub const CURVE_GY: [Chunk; NLEN] = [
    0x2BEC68505FE8632,
    0x5D5650CA0365DB1,
    0x3811C7EF435B6DB,
    0x7853D1B14B46C,
    0x56502E18E1C161D,
    0xC0DC616B,
];

pub const MODBYTES: usize = 42;
pub const BASEBITS: usize = 60;

pub const MODBITS: usize = 336;
pub const MOD8: usize = 5;
pub const MODTYPE: ModType = ModType::PseudoMersenne;
pub const SH: usize = 24;

pub const CURVETYPE: CurveType = CurveType::Edwards;
pub const CURVE_PAIRING_TYPE: CurvePairingType = CurvePairingType::Not;
pub const SEXTIC_TWIST: SexticTwist = SexticTwist::Not;
pub const ATE_BITS: usize = 0;
pub const SIGN_OF_X: SignOfX = SignOfX::Not;
pub const HASH_TYPE: usize = 48;
pub const AESKEY: usize = 24;
